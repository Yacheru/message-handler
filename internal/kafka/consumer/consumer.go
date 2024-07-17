package consumer

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	Topic []string
}

func NewConsumer(topic []string) *Consumer {
	return &Consumer{
		Topic: topic,
	}
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	go func() {
		select {
		case <-session.Context().Done():
			session.Commit()
		default:
			for msg := range claim.Messages() {

				// Место для обработки полученного сообщения

				session.MarkMessage(msg, "")
			}
		}
	}()

	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
