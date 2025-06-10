package producer

import (
	"context"
	"time"
)

type ProducerConfig struct {
	Type    string        // mq 类型: "kafka", "rabbitmq"
	Brokers []string      // broker 地址列表
	Timeout time.Duration // 超时设置
}

func NewMessageProducer(config ProducerConfig) Client {
	switch config.Type {
	case "kafka":
		return newKafkaProducer(config)
	default:
		panic("unknown mq type: " + config.Type)
	}
}

type Client interface {
	Send(ctx context.Context, destination string, message string) error
	Close() bool
}
