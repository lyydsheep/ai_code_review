package consumer

import (
	"github.com/IBM/sarama"
)

func NewConfig() *sarama.Config {
	// 配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// 手动提交
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	return config
}

func NewConsumerGroup(addrs []string, groupID string) sarama.ConsumerGroup {
	config := NewConfig()
	consumerGroup, err := sarama.NewConsumerGroup(addrs, groupID, config)
	if err != nil {
		panic(err)
	}
	return consumerGroup
}
