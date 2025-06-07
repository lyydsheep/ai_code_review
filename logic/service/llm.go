package service

import (
	"context"
	"github.com/faiz/llm-code-review/common/enum"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/logic/infrastructure/llm"
)

type LLMService struct {
	Client *llm.Client
}

func NewLLMService(client *llm.Client) *LLMService {
	return &LLMService{Client: client}
}

func (svc *LLMService) SetClient(ctx context.Context, category string) {
	switch category {
	case enum.DEEPSEEK:
		log.New(ctx).Info("set deepseek sender")
		svc.Client.SetSender(llm.NewDeepSeekSender())
	default:
		log.New(ctx).Info("set default sender", "category", category)
		svc.Client.SetSender(llm.NewDeepSeekSender())
	}

}

func (svc *LLMService) ListAllClient() []string {
	return []string{enum.DEEPSEEK}
}

func (svc *LLMService) GetReport(ctx context.Context, diffInfo string) (string, error) {
	return svc.Client.Sender.Send(ctx, diffInfo)
}
