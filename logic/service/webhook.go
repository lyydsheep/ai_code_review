package service

import (
	"context"
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/logic/infrastructure"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/producer"
	"github.com/faiz/llm-code-review/logic/repository"
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
	// 获取 diff 信息
	diff, err := svc.Github.GetDiff(ctx, user, *hook)
	if err != nil {
		return err
	}

	// 发送给消费者
	// TODO 根据用户类型 选择不同的队列
	// 实现优先级机制
	if err = svc.Kafka.Send(ctx, "anonymous-code-review", diff); err != nil {
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
