package consumer

import (
	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func NewConsumerGroup(ctx context.Context, topic []string) error {
	cfg := sarama.NewConfig()

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	consumerGroup, err := sarama.NewConsumerGroup(config.ServerConfig.KafkaBrokers, config.ServerConfig.KafkaConsumerGroup, cfg)
	if err != nil {
		logger.FatalF(err.Error(), logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

		return err
	}

	return Subscribe(ctx, topic, consumerGroup)
}

func Subscribe(ctx context.Context, topic []string, consumerGroup sarama.ConsumerGroup) error {
	consumer := NewConsumer(topic)

	go func() {
		logger.DebugF("starting consume messages...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})
		if err := consumerGroup.Consume(ctx, consumer.Topic, consumer); err != nil {
			logger.ErrorF("Error on consumer message: %v", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer}, err.Error())
		}
		if ctx.Err() != nil {
			return
		}
	}()

	return nil
}
