package service

import (
	"context"
	"encoding/json"
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/event"
	"github.com/faiz/llm-code-review/logic/infrastructure"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/producer"
	"github.com/faiz/llm-code-review/logic/repository"
	"github.com/google/uuid"
	"strings"
	"time"
)

type WebHookServiceV1 struct {
	UsrUserRepository repository.UsrUserRepository
	Github            infrastructure.GithubClient
	Kafka             producer.Client
}

func (svc *WebHookServiceV1) ProcessHook(ctx context.Context, hook *request.HookRequest) error {
	// 查表获取 user 信息
	user, err := svc.UsrUserRepository.GetUserByUsername(ctx, hook.UserName)
	if err != nil {
		return err
	}
	log.New(ctx).Debug("get user info", "user", user)
	// 获取 diff 信息
	diff, err := svc.Github.GetDiff(ctx, user, *hook)
	if err != nil {
		return err
	}
	log.New(ctx).Debug("get diff info", "diff", diff)

	// 发送给消费者
	// TODO 根据用户类型 选择不同的队列
	// 实现优先级机制
	id := strings.Replace(uuid.New().String(), "-", "", -1)
	push := event.Push{
		ID:         id,
		Priority:   event.LowPriority,
		Diff:       diff,
		Repository: hook.Repository,
		Time:       time.Now().UTC(),
		Username:   hook.UserName,
	}

	msg, err := json.Marshal(push)
	if err != nil {
		log.New(ctx).Error("Failed to marshal push: %v", "err", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("Failed to marshal push message")
	}
	if err = svc.Kafka.Send(ctx, "anonymous-code-review", msg); err != nil {
		log.New(ctx).Error("Failed to send message: %v", "err", err)
		return err
	}

	return nil
}

func NewWebHookServiceV1(usrUserRepository repository.UsrUserRepository,
	github infrastructure.GithubClient, kafka producer.Client) WebHookService {
	return &WebHookServiceV1{
		UsrUserRepository: usrUserRepository,
		Github:            github,
		Kafka:             kafka,
	}
}
