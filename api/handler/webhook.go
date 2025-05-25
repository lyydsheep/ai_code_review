package handler

import (
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/app"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	svc service.WebHookService
}

func NewWebhookHandler(svc service.WebHookService) WebhookHandler {
	return WebhookHandler{
		svc: svc,
	}
}

func (h *WebhookHandler) ProcessHook(ctx *gin.Context) {
	var req request.HookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.New(ctx).Error("Failed to bind JSON: %v", err)
		app.NewResponse(ctx).Error(err)
		return
	}
	// service 逻辑
	if err := h.svc.ProcessHook(ctx, &req); err != nil {
		log.New(ctx).Error("Failed to process hook: %v", err)
		app.NewResponse(ctx).Error(err)
		return
	}

	app.NewResponse(ctx).SuccessOk()
}
