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

	KafkaConsumerGroup string `mapstructure:"KAFKA_CONSUMER_GROUP"`
	KafkaBroker        string `mapstructure:"KAFKA_BROKER"`
	KafkaTopic         string `mapstructure:"KAFKA_TOPIC"`

	PSQLPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PSQLUser     string `mapstructure:"POSTGRES_USER"`
	PSQLDb       string `mapstructure:"POSTGRES_DB"`
	PSQLDsn      string `mapstructure:"POSTGRES_DSN"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
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
		ServerConfig.KafkaTopic == "" || ServerConfig.KafkaBroker == "" ||
		ServerConfig.PSQLPassword == "" || ServerConfig.PSQLDb == "" || ServerConfig.PSQLDsn == "" || ServerConfig.PSQLUser == "" {
		logger.Error("missing requirement variable!", logrus.Fields{constants.LoggerCategory: constants.Config})

		return constants.ErrMissVar
	}

	return nil
}
