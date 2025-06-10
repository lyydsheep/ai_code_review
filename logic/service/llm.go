package service

import (
	"github.com/faiz/llm-code-review/logic/infrastructure/llm"
)

type LLMService struct {
	Client *llm.Client
}

func NewLLMService(client *llm.Client) *LLMService {
	return &LLMService{Client: client}
}

// llm client
// consumer
