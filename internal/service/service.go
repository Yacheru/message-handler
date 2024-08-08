package service

import (
	"Messaggio/internal/repository"
	service "Messaggio/internal/service/messages"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"Messaggio/internal/entity"
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

type Services struct {
	Messages
}

func NewService(repo *repository.Repository) *Services {
	return &Services{
		Messages: service.NewMessagesService(repo.Messages),
	}
}
