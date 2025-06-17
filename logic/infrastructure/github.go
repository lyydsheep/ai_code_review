package infrastructure

import (
	"context"
	"fmt"
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util"
	"github.com/faiz/llm-code-review/common/util/httptool"
	"github.com/faiz/llm-code-review/dal/model"
)

type GithubClient interface {
	GetDiff(ctx context.Context, user model.UsrUser, hook request.HookRequest) (string, error)
}

type DefaultGithubClient struct {
}

func (client *DefaultGithubClient) GetDiff(ctx context.Context, user model.UsrUser, hook request.HookRequest) (string, error) {
	tokenBytes, err := util.DecryptAES(user.Token, []byte(user.AesKey))
	if err != nil {
		return "", errcode.ErrServer.WithCause(err).AppendMsg("Failed to decrypt token")
	}
	token := string(tokenBytes)

	// 获取 compare 信息
	respCode, respBody, err := httptool.Get(ctx, joinURL(hook.UserName, hook.Repository, hook.Compare),
		httptool.WithHeaders(map[string]string{
			"Accept":               "application/vnd.github.diff",
			"Authorization":        fmt.Sprintf("Bearer %s", token),
			"X-GitHub-Api-Version": "2022-11-28",
		}))
	if err != nil {
		log.New(ctx).Error("Failed to get compare: %v", err)
		return "", errcode.ErrServer.WithCause(err).AppendMsg("Failed to get compare")
	}
	if respCode != 200 {
		log.New(ctx).Error("Failed to get compare: %v", err)
		return "", errcode.ErrServer.WithCause(err).AppendMsg("Failed to get compare")
	}
	return string(respBody), nil
}

func joinURL(username, repository, compare string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/compare/%s", username, repository, compare)
}

func NewDefaultGithubClient() GithubClient {
	return &DefaultGithubClient{}
}
