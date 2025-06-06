// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
)

// Injectors from wire.go:

func InitializeApp() *gin.Engine {
	query := dao.NewQuery()
	cacheCache := cache.NewMemCache()
	usrUserRepository := repository.NewUsrUserRepositoryV1(query, cacheCache)
	webHookService := service.NewWebHookServiceV1(usrUserRepository)
	webhookHandler := handler.NewWebhookHandler(webHookService)
	v := middleware.GetHandlerFunc()
	engine := router.RegisterRoutersAndMiddleware(webhookHandler, v...)
	return engine
}
