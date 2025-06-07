package llm

import "context"

type Client struct {
	Sender Sender
}

// 使用前指定策略
func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetSender(sender Sender) {
	c.Sender = sender
}

type Sender interface {
	Send(context.Context, string) (string, error)
}

type SenderStrategy func(ctx context.Context, diffInfo string) (string, error)

func (s SenderStrategy) Send(ctx context.Context, diffInfo string) (string, error) {
	return s(ctx, diffInfo)
}
