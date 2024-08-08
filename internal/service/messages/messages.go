package service

import (
	"Messaggio/internal/entity"
	"Messaggio/internal/repository"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessagesService struct {
	messages repository.Messages
}

func NewMessagesService(messages repository.Messages) *MessagesService {
	return &MessagesService{messages}
}

func (s *MessagesService) InsertNew(ctx *gin.Context, message entity.Message) (*entity.DBMessage, error) {
	return s.messages.InsertNew(ctx, message)
}

func (s *MessagesService) GetAll(ctx *gin.Context) ([]entity.DBMessage, error) {
	return s.messages.GetAll(ctx)
}

func (s *MessagesService) GetById(ctx *gin.Context, id string) (entity.DBMessage, error) {
	return s.messages.GetById(ctx, id)
}

func (s *MessagesService) DeleteMessage(ctx *gin.Context, id string) (entity.DBMessage, error) {
	return s.messages.DeleteMessage(ctx, id)
}

func (s *MessagesService) EditMessage(ctx *gin.Context, id string, message entity.Message) (*entity.DBMessage, error) {
	return s.messages.EditMessage(ctx, id, message)
}

func (s *MessagesService) GetStats(ctx *gin.Context) (*entity.Statistic, error) {
	return s.messages.GetStats(ctx)
}

func (s *MessagesService) Mark(ctx context.Context, id uuid.UUID) error {
	return s.messages.Mark(ctx, id)
}
