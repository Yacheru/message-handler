package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/internal/http/routes"
	"Messaggio/internal/kafka/producer"
	"Messaggio/internal/repository"
	"Messaggio/pkg/constants"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(ctx context.Context) (*Server, error) {
	db, err := repository.NewPostgresConnection(ctx, &config.ServerConfig)

	KafkaProducer, err := producer.NewKafkaProducer([]string{config.ServerConfig.KafkaBroker}, []string{config.ServerConfig.KafkaTopic})
	if err != nil {
		return nil, err
	}

	router := setupRouter()

	api := router.Group("/messaggio")
	routes.NewRoute(ctx, api, KafkaProducer, db, []string{config.ServerConfig.KafkaTopic}).Routes()

	server := &http.Server{
		Addr:           ":" + config.ServerConfig.APIPort,
		Handler:        router,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1mb
	}

	return &Server{
		HTTPServer: server,
	}, nil
}

func (s *Server) Run() error {
	go func() {
		logger.Info("starting http server...", logrus.Fields{constants.LoggerCategory: constants.Server})
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Server})
		}
		logger.InfoF("server started on :%s port", logrus.Fields{constants.LoggerCategory: constants.Server}, config.ServerConfig.APIPort)
	}()

	return nil
}

func setupRouter() *gin.Engine {
	var mode = gin.ReleaseMode
	if config.ServerConfig.APIDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))

	return router
}
