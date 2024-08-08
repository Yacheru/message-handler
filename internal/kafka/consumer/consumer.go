package consumer

import (
	"Messaggio/init/logger"
	"Messaggio/internal/entity"
	"Messaggio/internal/service"
	"Messaggio/pkg/constants"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	Topic   []string
	service *service.Services
	ctx     context.Context
}

func NewKafkaConsumer(topic []string, service *service.Services, ctx context.Context) *Consumer {
	return &Consumer{
		Topic:   topic,
		service: service,
		ctx:     ctx,
	}
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var dbMessage entity.DBMessage

	for {
		select {
		case <-session.Context().Done():
			session.Commit()
		default:
			for msg := range claim.Messages() {
				logger.DebugF("claimed message: %v", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer}, string(msg.Value))

				err := json.Unmarshal(msg.Value, &dbMessage)
				if err != nil {
					return err
				}

				err = c.service.Mark(c.ctx, dbMessage.ID)
				if err != nil {
					return err
				}

				session.MarkMessage(msg, "mark message!")
			}
		}
	}
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
