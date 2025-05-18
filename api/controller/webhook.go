package controller

import (
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/app"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) ProcessHook(ctx *gin.Context) {
	var req request.HookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.New(ctx).Error("Failed to bind JSON: %v", err)
		app.NewResponse(ctx).Error(err)
		return
	}
	// service 逻辑

	app.NewResponse(ctx).SuccessOk()
}
