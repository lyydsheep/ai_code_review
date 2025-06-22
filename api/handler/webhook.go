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
		log.New(ctx).Error("failed to bind JSON", "err", err)
		app.NewResponse(ctx).Error(err)
		return
	}
	log.New(ctx).Debug("get hook request", "request", req)
	// service 逻辑
	if err := h.svc.ProcessHook(ctx, &req); err != nil {
		log.New(ctx).Error("failed to process hook.", "err", err)
		app.NewResponse(ctx).Error(err)
		return
	}
	log.New(ctx).Debug("process hook success")
	app.NewResponse(ctx).SuccessOk()
}
