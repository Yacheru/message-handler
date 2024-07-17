package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
)

var ServerConfig Config

type Config struct {
	APIPort  string `mapstructure:"API_PORT"`
	APIDebug bool   `mapstructure:"API_DEBUG"`

	KafkaConsumerGroup string   `mapstructure:"KAFKA_CONSUMER_GROUP"`
	KafkaBrokers       []string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopics        []string `mapstructure:"KAFKA_TOPICS"`

	PSQLHost     string `mapstructure:"POSTGRES_HOST"`
	PSQLPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PSQLUser     string `mapstructure:"POSTGRES_USER"`
	PSQLDb       string `mapstructure:"POSTGRES_DB"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.ErrorF("error reading config file, %v", logrus.Fields{constants.LoggerCategory: constants.Config}, err.Error())

		return err
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		logger.ErrorF("error unmarshal config, %v", logrus.Fields{constants.LoggerCategory: constants.Config}, err.Error())

		return err
	}

	if ServerConfig.APIPort == "" ||
		len(ServerConfig.KafkaTopics) == 0 || len(ServerConfig.KafkaBrokers) == 0 ||
		ServerConfig.PSQLPassword == "" || ServerConfig.PSQLHost == "" {
		logger.Error("missing requirement variable!", logrus.Fields{constants.LoggerCategory: constants.Config})

		return constants.ErrMissVar
	}

	return nil
}
