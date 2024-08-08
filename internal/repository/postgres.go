package repository

import (
	"Messaggio/init/config"
	"Messaggio/init/logger"
	"Messaggio/internal/entity"
	"Messaggio/pkg/constants"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

type Messages interface {
	InsertNew(ctx *gin.Context, message entity.Message) (*entity.DBMessage, error)
	GetAll(ctx *gin.Context) ([]entity.DBMessage, error)
	GetById(ctx *gin.Context, id string) (entity.DBMessage, error)
	DeleteMessage(ctx *gin.Context, id string) (entity.DBMessage, error)
	EditMessage(ctx *gin.Context, id string, message entity.Message) (*entity.DBMessage, error)
	GetStats(ctx *gin.Context) (*entity.Statistic, error)
	Mark(ctx context.Context, id uuid.UUID) error
}

type Repository struct {
	Messages
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewMessagesPostgres(db),
	}
}

func NewPostgresConnection(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	logger.Info("open postgresql connection...", logrus.Fields{constants.LoggerCategory: constants.Database})

	db, err := sqlx.Open("postgres", cfg.PSQLDsn)
	if err != nil {
		return nil, err
	}

	logger.Info("successful connection to postgres. Migrating...", logrus.Fields{constants.LoggerCategory: constants.Database})

	if err := goose.UpContext(ctx, db.DB, "./schema"); err != nil {
		logger.ErrorF("error using migrates: %v", logrus.Fields{constants.LoggerCategory: constants.Database}, err)

		return nil, err
	}

	return db, err
}
