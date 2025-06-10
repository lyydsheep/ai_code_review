package test

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/faiz/llm-code-review/logic/infrastructure/mq/producer"
	"log"
	"sync"
	"testing"
)

func TestProduce(t *testing.T) {
	svc := producer.NewMessageProducer(producer.ProducerConfig{
		Type:    "kafka",
		Brokers: []string{":9092"},
	})
	err := svc.Send(context.Background(), "topic-test", "hello world")
	if err != nil {
		log.Fatal(err)
	}
}

func TestConsume(t *testing.T) {
	// 配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}

	// 创建消费者组
	group := "test-group"
	broker := []string{":9092"}
	topic := "topic-test"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cg, err := sarama.NewConsumerGroup(broker, group, config)
	if err != nil {
		log.Fatalf("failed to start consumer group: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Printf("start consume")
			if err := cg.Consume(ctx, []string{topic}, &consumerGroupHandler{}); err != nil {
				log.Fatalf("failed to start consumer group: %v", err)
			}
		}
	}()
	wg.Wait()
}

type consumerGroupHandler struct{}

func (h *consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("setup")
	return nil
}

func (h *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("cleanup")
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	for msg := range msgs {
		log.Printf("topic is %s, value is %s", msg.Topic, msg.Value)
		// 标记消息，并手动提交
		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}
