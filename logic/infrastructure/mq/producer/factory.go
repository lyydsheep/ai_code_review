package producer

import (
	"context"
	"github.com/faiz/llm-code-review/config"
	"time"
)

type Config struct {
	Type    string        // mq 类型: "kafka", "rabbitmq"
	Brokers []string      // broker 地址列表
	Timeout time.Duration // 超时设置
}

func NewKafkaConfig() Config {
	return Config{
		Type:    "kafka",
		Brokers: config.Kafka.Brokers,
		Timeout: config.Kafka.Timeout,
	}
}

func NewMessageProducer(config Config) Client {
	switch config.Type {
	case "kafka":
		return newKafkaProducer(config)
	default:
		panic("unknown mq type: " + config.Type)
	}
}

type Client interface {
	Send(ctx context.Context, destination string, message []byte) error
	Close() bool
}
