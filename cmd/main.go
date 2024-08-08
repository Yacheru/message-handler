package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"

	"Messaggio/cmd/server"
	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Config})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.Config})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	app, err := server.NewServer(ctx)
	if err != nil {
		logger.Error("error create a new http-server", logrus.Fields{constants.LoggerCategory: constants.Server})

		cancel()
	}
	if err = app.Run(); err != nil {
		logger.Fatal("error run http-server", logrus.Fields{constants.LoggerCategory: constants.Server})

		cancel()
	}

	<-ctx.Done()

	logger.Info("server shutdown", logrus.Fields{constants.LoggerCategory: constants.Server})
}
