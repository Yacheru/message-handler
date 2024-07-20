package consumer

import (
	"Messaggio/internal/repository"
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
)

func NewConsumerGroup(ctx context.Context, topic []string, postgres *repository.Postgres) error {
	logger.Info("Create a new consumer group...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

	cfg := sarama.NewConfig()

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	consumerGroup, err := sarama.NewConsumerGroup([]string{config.ServerConfig.KafkaBroker}, config.ServerConfig.KafkaConsumerGroup, cfg)
	if err != nil {
		logger.ErrorF(err.Error(), logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

		return err
	}

	return Subscribe(ctx, topic, postgres, consumerGroup)
}

func Subscribe(ctx context.Context, topic []string, postgres *repository.Postgres, consumerGroup sarama.ConsumerGroup) error {
	consumer := NewKafkaConsumer(topic, postgres, ctx)

	logger.Info("Starting consumer session...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, consumer.Topic, consumer); err != nil {
				logger.ErrorF("Error start consumer session: %v", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer}, err.Error())
			}
		}
	}()

	return nil
}
