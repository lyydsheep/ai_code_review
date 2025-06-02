package test

import (
	"context"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq"
	"log"
	"testing"
)

func TestProduce(t *testing.T) {
	svc := mq.NewMessageProducer(mq.ProducerConfig{
		Type:    "kafka",
		Brokers: []string{"47.120.11.159:9092"},
	})
	err := svc.Send(context.Background(), "topic-test", "hello world")
	if err != nil {
		log.Fatal(err)
	}
}
