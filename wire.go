//go:build wireinject

package main

import (
	"github.com/faiz/llm-code-review/api/handler"
	"github.com/faiz/llm-code-review/api/router"
	"github.com/faiz/llm-code-review/common/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, handler.NewBuildController,
	)
	return nil
}
