package llm

import (
	"context"
	"github.com/faiz/llm-code-review/common/enum"
	log "github.com/faiz/llm-code-review/common/logger"
)

type Client struct {
	Sender Sender
}

// 使用前指定策略
func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetSender(ctx context.Context, category string) {
	switch category {
	case enum.DEEPSEEK:
		log.New(ctx).Info("set deepseek sender")
		c.Sender = NewDeepSeekSender()
	default:
		log.New(ctx).Info("set default sender", "category", category)
		c.Sender = NewDeepSeekSender()
	}

}

type Sender interface {
	Send(context.Context, string) (string, error)
}

type SenderStrategy func(ctx context.Context, diffInfo string) (string, error)

func (s SenderStrategy) Send(ctx context.Context, diffInfo string) (string, error) {
	return s(ctx, diffInfo)
}

func (c *Client) ListAllClient() []string {
	return []string{enum.DEEPSEEK}
}
