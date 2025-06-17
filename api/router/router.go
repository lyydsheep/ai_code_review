package router

// 路由注册相关的, 都放在router 包下

import (
	"github.com/faiz/llm-code-review/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersAndMiddleware(webhook handler.WebhookHandler, fs ...gin.HandlerFunc) *gin.Engine {
	s := gin.Default()
	s.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    404,
			"message": "The requested API does not exist",
		})
	})
	s.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(405, gin.H{
			"code":    405,
			"message": "The requested API does not support the request method",
		})
	})
	RegisterMiddleware(s, fs...)

	g := s.Group("/")
	registerWebhook(g, webhook)
	return s
}

func RegisterMiddleware(g *gin.Engine, fs ...gin.HandlerFunc) *gin.Engine {
	g.Use(fs...)
	return g
}
