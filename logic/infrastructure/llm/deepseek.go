package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util/httptool"
	"net/http"
	"os"
	"time"
)

const (
	deepseekAPI = "https://api.deepseek.com/chat/completions"
)

type DeepseekRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepseekResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

// NewDeepseekStrategy 创建一个新的 Deepseek 调用策略
func newDeepSeekSender() SenderStrategy {
	return func(ctx context.Context, diffInfo string) (string, error) {
		req := DeepseekRequest{
			Model: "deepseek-chat",
			Messages: []Message{
				{
					Role:    "system",
					Content: systemPrompt,
				},
				{
					Role:    "user",
					Content: userPrompt + diffInfo,
				},
			},
			Stream: false,
		}

		jsonData, err := json.Marshal(req)
		if err != nil {
			log.New(ctx).Error("marshal request failed: %w", err)
			return "", errcode.ErrServer.WithCause(err).AppendMsg("marshal request failed")
		}

		resCode, body, err := httptool.Post(ctx, jsonData, deepseekAPI,
			httptool.WithAuthorization(os.Getenv("DEEPSEEK_API_KEY")),
			httptool.WithTimeout(time.Minute*30))

		if resCode != http.StatusOK {
			log.New(ctx).Info("http code is not 200", "http code", resCode, "body", string(body))
			return "", errcode.ErrServer.WithCause(err).AppendMsg("http code is not 200")
		}

		var deepseekResp DeepseekResponse
		// 流式处理，适合大对象
		if err = json.NewDecoder(bytes.NewReader(body)).Decode(&deepseekResp); err != nil {
			log.New(ctx).Error("decode response failed", "err", err)
			return "", errcode.ErrServer.WithCause(err).AppendMsg("decode response failed")
		}

		if len(deepseekResp.Choices) == 0 {
			log.New(ctx).Error("no response content")
			return "", errcode.ErrServer.WithCause(err).AppendMsg("no response content")
		}

		log.New(ctx).Info("deepseek response", "content", deepseekResp.Choices[0].Message.Content)
		return deepseekResp.Choices[0].Message.Content, nil
	}
}
