package handlers

import (
	"Messaggio/init/logger"
	"Messaggio/internal/entity"
	"Messaggio/internal/kafka/producer"
	"Messaggio/internal/repository"
	"Messaggio/pkg/constants"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	producer *producer.Producer
	postgres *repository.Postgres
}

func NewHandlers(producer *producer.Producer, postgres *repository.Postgres) *Handlers {
	return &Handlers{
		producer: producer,
		postgres: postgres,
	}
}

func (h *Handlers) InsertNew(ctx *gin.Context) {
	var message entity.Message
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Your message is invalid")
		return
	}

	dbMessage, err := h.postgres.InsertNew(ctx, message)
	if err != nil {
		logger.ErrorF("Error insert message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	msg, err := json.Marshal(dbMessage)
	if err != nil {
		logger.ErrorF("Error marshal message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.producer.SendMessage(msg)

	NewSuccessResponse(ctx, http.StatusOK, "Message added successfully", dbMessage)
}

func (h *Handlers) GetAll(ctx *gin.Context) {
	dbMessages, err := h.postgres.GetAll(ctx)
	logger.InfoF("%v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, dbMessages)

	if err != nil {
		logger.ErrorF("Error get all messages: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Here all messages!", dbMessages)
}

func (h *Handlers) GetByID(ctx *gin.Context) {
	dbMessage, err := h.postgres.GetById(ctx, ctx.Param("id"))
	logger.InfoF("%v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, dbMessage)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(ctx, http.StatusInternalServerError, "Message does not exist")
			return
		default:
			logger.ErrorF("Error get message by id: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(ctx, http.StatusOK, "Get By ID!", dbMessage)
}

func (h *Handlers) DeleteMessage(ctx *gin.Context) {
	dbMessage, err := h.postgres.DeleteMessage(ctx, ctx.Param("id"))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(ctx, http.StatusInternalServerError, "Message does not exist")
			return
		default:
			logger.ErrorF("Error delete message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(ctx, http.StatusOK, "Delete Message!", dbMessage)
}

func (h *Handlers) EditMessage(ctx *gin.Context) {
	var message entity.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "your message is invalid")
		return
	}

	dbMessage, err := h.postgres.EditMessage(ctx, ctx.Param("id"), message)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(ctx, http.StatusNotFound, "Message does not exist")
			return
		default:
			logger.ErrorF("Error edit message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(ctx, http.StatusOK, "Edit Message!", dbMessage)
}

func (h *Handlers) GetStats(ctx *gin.Context) {
	stats, err := h.postgres.GetStats(ctx)
	if err != nil {
		logger.InfoF("Error get statistic: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Here stats", stats)
}
