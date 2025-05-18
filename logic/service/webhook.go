package appService

import (
	"context"
	"github.com/faiz/llm-code-review/api/request"
)

type WebHookServiceV1 struct{}

func (svc *WebHookServiceV1) ProcessHook(ctx context.Context, hook *request.HookRequest) error {
	// TODO
	panic("")
}

func NewWebHookServiceV1() WebHookService {
	return &WebHookServiceV1{}
}
