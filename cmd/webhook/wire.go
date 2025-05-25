//go:build wireinject

package main

import (
	"github.com/faiz/llm-code-review/api/handler"
	"github.com/faiz/llm-code-review/api/router"
	"github.com/faiz/llm-code-review/common/middleware"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/logic/repository"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, handler.NewWebhookHandler,
		service.NewWebHookServiceV1, repository.NewUsrUserRepositoryV1,
		cache.NewMemCache, dao.NewQuery)

	return nil
}
