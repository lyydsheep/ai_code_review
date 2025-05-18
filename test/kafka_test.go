package test

import (
	"github.com/IBM/sarama"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{"47.120.11.159:9092"}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))

		msg := &sarama.ProducerMessage{Topic: "topic-test", Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}
