package consumer

import (
	"github.com/IBM/sarama"
	"log"
)

type PriorityHandler struct {
	messages chan *sarama.ConsumerMessage
	finished chan struct{}
	consumerGroupHandler
}

func (h *PriorityHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// 将消息发送至指定优先级的 channel 中
		h.messages <- message
		// 等待消息处理完成
		<-h.finished
		session.MarkMessage(message, "")
		session.Commit()
	}
	return nil
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
	log.Println("consumeClaim")
	return nil
}
