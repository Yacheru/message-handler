package repository

import (
	"Messaggio/init/logger"
	"Messaggio/pkg/constants"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewPostgresConnection(dsn string) (*sqlx.DB, error) {
	logger.Debug("Open postgresql connection...", logrus.Fields{constants.LoggerCategory: constants.Database})

	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, err
}
