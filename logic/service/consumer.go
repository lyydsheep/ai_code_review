package service

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/config"
	"github.com/faiz/llm-code-review/event"
	"github.com/faiz/llm-code-review/logic/infrastructure/llm"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/consumer"
	"golang.org/x/sync/errgroup"
	"time"
)

type Consumer struct {
	LLMClient  llm.Sender
	MQHandlers []*consumer.PriorityHandler
	Channels   map[string]chan *sarama.ConsumerMessage
	Finishes   map[string]chan struct{}
}

func NewLLMService(client llm.Sender, mqHandlers ...*consumer.PriorityHandler) *Consumer {
	channels := make(map[string]chan *sarama.ConsumerMessage)
	finish := make(map[string]chan struct{})
	for _, handler := range mqHandlers {
		channels[handler.Topic] = handler.Messages
		finish[handler.Topic] = handler.Finished
	}
	return &Consumer{LLMClient: client, MQHandlers: mqHandlers, Channels: channels}
}

func (svc *Consumer) Run(ctx context.Context) {
	log.New(ctx).Info("starting consumer")
	startChannels(ctx, svc.MQHandlers)
	startConsumers(ctx, svc.Channels, svc.Finishes)
}

func startConsumers(ctx context.Context, messages map[string]chan *sarama.ConsumerMessage, finish map[string]chan struct{}) {
	log.New(ctx).Info("starting consumer service")
	for {
		// 依据时间片对不同的handler 进行处理
		select {
		case msg := <-messages[event.HighPriority]:
			log.New(ctx).Debug("received message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "priority", event.HighPriority)
			consumeEvent(ctx, msg, finish[event.LowPriority])
		case <-time.After(750 * time.Millisecond):
			// 下放
			select {
			case msg := <-messages[event.LowPriority]:
				log.New(ctx).Debug("received message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "priority", event.HighPriority)
				consumeEvent(ctx, msg, finish[event.LowPriority])
			case <-time.After(250 * time.Millisecond):
			}
		}
	}
}

func consumeEvent(ctx context.Context, msg *sarama.ConsumerMessage, finish chan struct{}) {
	// TODO 待实现
	// 接受消息
	// 插入数据库（幂等性校验）, 提交 commit。 实际消费动作只有插入数据操作
	// 向 llm 发送请求，获取报告
	// 发送邮件
}

func startChannels(ctx context.Context, handlers []*consumer.PriorityHandler) {
	log.New(ctx).Info("starting consumer channels")
	// 获取  consumer group
	cg := consumer.NewConsumerGroup(config.Kafka.Brokers, "llm-code-review-consumer-group")

	// 启动消费接收器
	var eg errgroup.Group
	for _, handler := range handlers {
		log.New(ctx).Debug("starting consumer", "topic", handler.Topic)
		eg.Go(func() error { return consume(ctx, cg, handler) })
	}
	if err := eg.Wait(); err != nil {
		log.New(ctx).Fatal("consume error: %v", err)
	}
}

func consume(ctx context.Context, cg sarama.ConsumerGroup, handler *consumer.PriorityHandler) error {
	log.New(ctx).Info("starting consumer group", "topic", handler.Topic)
	if err := cg.Consume(ctx, []string{handler.Topic}, handler); err != nil {
		log.New(ctx).Error("consume error: %v", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("consume error")
	}
	return nil
}
