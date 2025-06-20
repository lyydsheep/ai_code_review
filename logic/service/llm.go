package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"github.com/faiz/llm-code-review/common/enum"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/config"
	"github.com/faiz/llm-code-review/dal/model"
	"github.com/faiz/llm-code-review/event"
	"github.com/faiz/llm-code-review/logic/infrastructure/llm"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/consumer"
	"github.com/faiz/llm-code-review/logic/repository"
	"github.com/gomarkdown/markdown"
	"golang.org/x/sync/errgroup"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"time"
)

type LLMService struct {
	LLMClient    *llm.Client
	MQHandlers   []*consumer.PriorityHandler
	Channels     map[string]chan *sarama.ConsumerMessage
	Finishes     map[string]chan struct{}
	PushInfoRepo repository.PushInfoRepository
	UserRepo     repository.UsrUserRepository
}

// NewLLMService topic必须是 high、middle、low 中的一个
func NewLLMService(client *llm.Client, repo repository.PushInfoRepository, userRepo repository.UsrUserRepository, mqHandlers ...*consumer.PriorityHandler) *LLMService {
	channels := make(map[string]chan *sarama.ConsumerMessage)
	finish := make(map[string]chan struct{})
	for _, handler := range mqHandlers {
		channels[handler.Topic] = handler.Messages
		finish[handler.Topic] = handler.Finished
	}
	return &LLMService{LLMClient: client, MQHandlers: mqHandlers, Channels: channels, Finishes: finish, PushInfoRepo: repo, UserRepo: userRepo}
}

func (svc *LLMService) Run(ctx context.Context) {
	log.New(ctx).Info("starting consumer")
	go svc.startChannels(ctx, svc.MQHandlers)
	go svc.startConsumers(ctx, svc.Channels, svc.Finishes)
}

// 添加一个Stop方法用于清理资源
func (svc *LLMService) Stop() {
	for topic, finish := range svc.Finishes {
		select {
		case finish <- struct{}{}:
			log.New(context.Background()).Info("sent finish signal", "topic", topic)
		default:
			log.New(context.Background()).Info("finish channel already closed", "topic", topic)
		}
	}
}

func (svc *LLMService) startConsumers(ctx context.Context, messages map[string]chan *sarama.ConsumerMessage, finish map[string]chan struct{}) {
	log.New(ctx).Info("starting consumer service")
	for {
		// 依据时间片对不同的handler 进行处理
		// 避免低优先级饿死
		ch := time.After(750 * time.Millisecond)
		select {
		case msg := <-messages[event.HighPriority]:
			log.New(ctx).Debug("received message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "priority", event.HighPriority)
			svc.consumeEvent(ctx, msg, finish[event.HighPriority])
		case <-ch:
			// 下放
			ch = time.After(250 * time.Millisecond)
			select {
			case msg := <-messages[event.LowPriority]:
				log.New(ctx).Debug("received message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "priority", event.LowPriority)
				svc.consumeEvent(ctx, msg, finish[event.LowPriority])
			case <-ch:
			}
		}
	}
}

func (svc *LLMService) consumeEvent(ctx context.Context, msg *sarama.ConsumerMessage, finish chan struct{}) {
	log.New(ctx).Info("consuming message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)
	// 接受消息
	var pushEvent event.Push
	if err := json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&pushEvent); err != nil {
		log.New(ctx).Error("decode message failed", "err", err)
		return
	}

	log.New(ctx).Debug("received message", "pushEvent", pushEvent)

	// 插入数据库（幂等性校验）, 提交 commit。 实际消费动作只有插入数据操作
	pushInfo, err := svc.PushInfoRepo.Create(ctx, eventToModel(pushEvent))
	if err != nil {
		// TODO 存入本地 + 异步重试
		finish <- struct{}{}
		log.New(ctx).Error("insert message failed", "pushEvent", pushEvent, "err", err)
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return
		}
		log.New(ctx).Info("message already exists", "pushEvent", pushEvent)

		if pushInfo.Status != model.Init {
			// 已经处理过且成功或失败的消息直接返回
			return
		}
		log.New(ctx).Info("message status is init. try to send again")
	}
	// 完成消费
	finish <- struct{}{}
	log.New(ctx).Debug("insert message success", "pushEvent", pushEvent)

	// 向 llm 发送请求，获取报告
	if err = svc.LLMClient.SetSender(ctx, enum.DEEPSEEK); err != nil {
		log.New(ctx).Error("set sender failed", "err", err)
		return
	}
	report, err := svc.LLMClient.Send(ctx, pushEvent.Diff)
	if err != nil {
		log.New(ctx).Error("send message to llm failed", "err", err)
		return
	}
	log.New(ctx).Debug("send message to llm success", "report", report)

	if report == "" {
		log.New(ctx).Info("no report generated.use the default text instead")
		// TODO 做一个 replace 操作
	}
	// 将 markdown 格式转换为 html 格式，并准备发送邮件
	html := string(markdown.ToHTML([]byte(report), nil, nil))

	user, err := svc.UserRepo.GetUserByUsername(ctx, pushInfo.Username)
	if err != nil {
		log.New(ctx).Error("get user failed", "err", err)
		return
	}
	log.New(ctx).Debug("get user info", "user", user)
	sendMsg(user.Email, html)
}

func sendMsg(to, content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Email.Username)
	m.SetHeader("Subject", "Your Code Review Report")
	m.SetHeader("To", to)
	m.SetBody("text/html", content)
	ch <- m
}

func eventToModel(push event.Push) model.PushInfo {
	return model.PushInfo{
		EventID:    push.ID,
		Username:   push.Username,
		Repository: push.Repository,
		Diff:       push.Diff,
		EventTime:  push.Time,
	}
}

func (svc *LLMService) startChannels(ctx context.Context, handlers []*consumer.PriorityHandler) {
	log.New(ctx).Info("starting consumer channels")
	// 获取  consumer group
	cg := consumer.NewConsumerGroup(config.Kafka.Brokers, "llm-code-review-consumer-group")

	// 启动消费接收器
	var eg errgroup.Group
	for _, handler := range handlers {
		log.New(ctx).Debug("starting consumer", "topic", handler.Topic)
		eg.Go(func() error { return svc.consume(ctx, cg, handler) })
	}
	if err := eg.Wait(); err != nil {
		log.New(ctx).Error("consume error: %v", err)
		// 在发生错误时停止所有资源
		svc.Stop()
	}
}

func (svc *LLMService) consume(ctx context.Context, cg sarama.ConsumerGroup, handler *consumer.PriorityHandler) error {
	log.New(ctx).Info("starting consumer", "topic", handler.Topic)
	if err := cg.Consume(ctx, []string{handler.Topic}, handler); err != nil {
		log.New(ctx).Error("consume error: %v", err)
		return err
	}
	return nil
}
