//go:build wireinject

package main

import (
	"github.com/faiz/llm-code-review/api/handler"
	"github.com/faiz/llm-code-review/api/router"
	"github.com/faiz/llm-code-review/common/middleware"
	"github.com/faiz/llm-code-review/config"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/logic/infrastructure"
	"github.com/faiz/llm-code-review/logic/infrastructure/llm"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/consumer"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/producer"
	"github.com/faiz/llm-code-review/logic/repository"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/google/wire"
)

func InitializeApp() *App {
	wire.Build(NewApp,
		router.RegisterRoutersAndMiddleware, consumer.NewMQHandlers, service.NewLLMService,
		middleware.GetHandlerFunc, handler.NewWebhookHandler, llm.NewClient,
		service.NewWebHookServiceV1, repository.NewUsrUserRepositoryV1, repository.NewPushInfoRepositoryV1,
		infrastructure.NewDefaultGithubClient, producer.NewMessageProducer, wire.Struct(new(producer.Config), "*"),
		wire.Value("kafka"), wire.Value(config.Kafka.Brokers), wire.Value(config.Kafka.Timeout),
		cache.NewMemCache, dao.NewQuery)

	return nil
}
