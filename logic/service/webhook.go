package service

import (
	"context"
	"fmt"
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util/httptool"
)

type WebHookServiceV1 struct{}

func (svc *WebHookServiceV1) ProcessHook(ctx context.Context, hook *request.HookRequest) error {
	// TODO
	// 查表获取 token 信息

	// 获取 compare 信息
	respCode, respBody, err := httptool.Get(ctx, joinURL(hook.UserName, hook.Repository, hook.Compare),
		httptool.WithHeaders(map[string]string{
			"Accept":               "application/vnd.github.diff",
			"X-GitHub-Api-Version": "2022-11-28",
		}))
	if err != nil {
		log.New(ctx).Error("Failed to get compare: %v", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("Failed to get compare")
	}
	if respCode != 200 {
		log.New(ctx).Error("Failed to get compare: %v", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("Failed to get compare")
	}

	// 放入 mq 中

	return nil
}

func joinURL(username, repository, compare string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/compare/%s", username, repository, compare)
}

func NewWebHookServiceV1() WebHookService {
	return &WebHookServiceV1{}
}
