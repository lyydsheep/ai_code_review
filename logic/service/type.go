package service

import (
	"context"
	"github.com/faiz/llm-code-review/api/request"
)

type WebHookService interface {
	ProcessHook(ctx context.Context, hook *request.HookRequest) (string, error)
}

type MQService interface {
	Send(ctx context.Context, destination string, message string) error
	Close() bool
}
