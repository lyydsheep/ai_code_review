package factory

import (
	"github.com/faiz/llm-code-review/logic/service"
	"time"
)

type ProducerConfig struct {
	Type    string        // mq 类型: "kafka", "rabbitmq"
	Brokers []string      // broker 地址列表
	Timeout time.Duration // 超时设置
}

func NewMessageProducer(config ProducerConfig) service.MQService {
	switch config.Type {
	case "kafka":
		return newKafkaProducer(config)
	default:
		panic("unknown mq type: " + config.Type)
	}
}
