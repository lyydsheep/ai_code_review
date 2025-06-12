package llm

import (
	"context"
	"fmt"
	"github.com/faiz/llm-code-review/common/enum"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
)

// Client 使用策略模式
// 在Send之前需要指定策略
type Client struct {
	sender Sender
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetSender(ctx context.Context, category string) error {
	switch category {
	case enum.DEEPSEEK:
		log.New(ctx).Info("set deepseek sender")
		c.sender = NewDeepSeekSender()
	default:
		err := fmt.Errorf("unsupported sender category: %s", category)
		log.New(ctx).Error("failed to set sender", "error", err)
		return err
	}
	return nil
}

func (c *Client) Send(ctx context.Context, diffInfo string) (string, error) {
	if c.sender == nil {
		log.New(ctx).Error("sender is nil")
		return "", errcode.ErrServer.AppendMsg("sender is nil")
	}
	return c.sender.Send(ctx, diffInfo)
}

type Sender interface {
	Send(context.Context, string) (string, error)
}

type SenderStrategy func(ctx context.Context, diffInfo string) (string, error)

func (s SenderStrategy) Send(ctx context.Context, diffInfo string) (string, error) {
	return s(ctx, diffInfo)
}

// ListAllClient 获取所有支持的模型
func (c *Client) ListAllClient() []string {
	return []string{enum.DEEPSEEK}
}
