package producer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"time"
)

type KafkaProducer struct {
	Client sarama.Client
}

func (k *KafkaProducer) Send(ctx context.Context, destination string, message string) error {
	producer, err := sarama.NewSyncProducerFromClient(k.Client)
	if err != nil {
		log.New(ctx).Error("Failed to create producer: %v", "err", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("Failed to create producer")
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: destination,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.New(ctx).Error("Failed to send message: %v", "err", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("Failed to send message")
	}
	log.New(ctx).Info("Message sent to partition %d with offset %d", "partition", partition, "offset", offset)
	return nil
}

func (k *KafkaProducer) Close() bool {
	return k.Client.Closed()
}

type Option interface {
	apply(*sarama.Config)
}

type fn func(*sarama.Config)

func (f fn) apply(config *sarama.Config) {
	f(config)
}

func WithTimeout(timeout time.Duration) Option {
	return fn(func(conf *sarama.Config) {
		conf.Producer.Timeout = timeout
	})
}

func WithReturnSuccess() Option {
	return fn(func(conf *sarama.Config) {
		conf.Producer.Return.Successes = true
	})
}

func WithWaitForAll() Option {
	return fn(func(conf *sarama.Config) {
		conf.Producer.RequiredAcks = sarama.WaitForAll
	})
}

func newConfig(options ...Option) *sarama.Config {
	conf := sarama.NewConfig()
	for _, opt := range options {
		opt.apply(conf)
	}
	return conf
}

func newKafkaProducer(proConf ProducerConfig) Client {
	if proConf.Timeout <= 0 {
		proConf.Timeout = 10 * time.Second
	}
	config := newConfig(WithTimeout(proConf.Timeout), WithReturnSuccess(), WithWaitForAll())
	client, err := sarama.NewClient(proConf.Brokers, config)
	if err != nil {
		panic(err)
	}
	return &KafkaProducer{
		Client: client,
	}
}
