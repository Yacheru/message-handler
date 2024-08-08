package producer

import (
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	Topics   []string
	Producer sarama.AsyncProducer
}

func NewKafkaProducer(brokers []string, topics []string) (*Producer, error) {
	logger.Info("create a new async producer...", logrus.Fields{constants.LoggerCategory: constants.KafkaConsumer})

	cfg := sarama.NewConfig()

	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		logger.ErrorF("failed create new producer: %v", logrus.Fields{constants.LoggerCategory: constants.KafkaProducer}, err.Error())

		return nil, err
	}

	return &Producer{
		Producer: producer,
		Topics:   topics,
	}, nil
}

func (p *Producer) SendMessage(message []byte) {
	p.Producer.Input() <- &sarama.ProducerMessage{
		Topic: p.Topics[0],
		Value: sarama.ByteEncoder(message),
	}

	logger.DebugF("%v sent to topic %v", logrus.Fields{constants.LoggerCategory: constants.KafkaProducer}, string(message), p.Topics[0])
}
