package consumer

import (
	"Messaggio/internal/service"
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
)

func NewConsumerGroup(ctx context.Context, topic []string, service *service.Services) error {
	logger.Info("create a new consumer group...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

	cfg := sarama.NewConfig()

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	consumerGroup, err := sarama.NewConsumerGroup([]string{config.ServerConfig.KafkaBroker}, config.ServerConfig.KafkaConsumerGroup, cfg)
	if err != nil {
		logger.ErrorF(err.Error(), logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

		return err
	}

	return Subscribe(ctx, topic, service, consumerGroup)
}

func Subscribe(ctx context.Context, topic []string, service *service.Services, consumerGroup sarama.ConsumerGroup) error {
	consumer := NewKafkaConsumer(topic, service, ctx)

	logger.Info("starting consumer session...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

	go func() {
		if err := consumerGroup.Consume(ctx, consumer.Topic, consumer); err != nil {
			logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})
		}
		if ctx.Err() != nil {
			return
		}
	}()

	return nil
}
