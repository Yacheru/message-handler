package main

import (
	"github.com/sirupsen/logrus"

	"Messaggio/cmd/server"
	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Config})
	}
	logger.Info("Configuration loaded", logrus.Fields{constants.LoggerCategory: constants.Config})
}

func main() {
	app, err := server.NewServer()
	if err != nil {
		logger.Fatal("Error create a new http-server", logrus.Fields{constants.LoggerCategory: constants.Server})
	}
	if err = app.Run(); err != nil {
		logger.Fatal("Error run http-server", logrus.Fields{constants.LoggerCategory: constants.Server})
	}
}
