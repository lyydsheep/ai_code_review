package appService

import (
	"context"
	"github.com/faiz/llm-code-review/api/request"
)

type WebHookService interface {
	ProcessHook(ctx context.Context, hook *request.HookRequest) error
}
