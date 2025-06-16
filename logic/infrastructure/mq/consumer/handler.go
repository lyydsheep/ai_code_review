package consumer

import (
	"github.com/IBM/sarama"
	"github.com/faiz/llm-code-review/event"
	"log"
)

type PriorityHandler struct {
	Topic    string
	Messages chan *sarama.ConsumerMessage
	Finished chan struct{}
	consumerGroupHandler
}

func NewPriorityHandler(topic string) *PriorityHandler {
	return &PriorityHandler{
		Topic:    topic,
		Messages: make(chan *sarama.ConsumerMessage),
		Finished: make(chan struct{}),
	}
}

// 目前只支持低、高优先级
func NewMQHandlers() []*PriorityHandler {
	handlers := make([]*PriorityHandler, 0)
	for _, topic := range []string{event.LowPriority, event.HighPriority} {
		handlers = append(handlers, NewPriorityHandler(topic))
	}
	return handlers
}

func (h *PriorityHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// 将消息发送至指定优先级的 channel 中
		h.Messages <- message
		// 等待消息处理完成
		<-h.Finished
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
