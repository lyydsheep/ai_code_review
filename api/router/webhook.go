package router

import (
	"github.com/faiz/llm-code-review/api/handler"
	"github.com/gin-gonic/gin"
)

func registerWebhook(s *gin.RouterGroup, webhook handler.WebhookHandler) {
	g := s.Group("/webhook")
	g.POST("/event", webhook.ProcessHook)
}
