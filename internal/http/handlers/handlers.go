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

func (h *Handlers) InsertNew(c *gin.Context) {
	var message entity.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Your message is invalid")
		return
	}

	dbMessage, err := h.postgres.InsertNew(c, message)
	if err != nil {
		logger.ErrorF("Error insert message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	msg, err := json.Marshal(dbMessage)
	if err != nil {
		logger.ErrorF("Error marshal message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.producer.SendMessage(msg)

	NewSuccessResponse(c, http.StatusOK, "Message added successfully", dbMessage)
}

func (h *Handlers) GetAll(c *gin.Context) {
	dbMessages, err := h.postgres.GetAll(c)
	logger.InfoF("%v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, dbMessages)

	if err != nil {
		logger.ErrorF("Error get all messages: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Here all messages!", dbMessages)
}

func (h *Handlers) GetByID(c *gin.Context) {
	dbMessage, err := h.postgres.GetById(c, c.Param("id"))
	logger.InfoF("%v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, dbMessage)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(c, http.StatusInternalServerError, "Message does not exist")
			return
		default:
			logger.ErrorF("Error get message by id: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(c, http.StatusOK, "Get By ID!", dbMessage)
}

func (h *Handlers) DeleteMessage(c *gin.Context) {
	dbMessage, err := h.postgres.DeleteMessage(c, c.Param("id"))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(c, http.StatusInternalServerError, "Message does not exist")
			return
		default:
			logger.ErrorF("Error delete message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(c, http.StatusOK, "Delete Message!", dbMessage)
}

func (h *Handlers) EditMessage(c *gin.Context) {
	var message entity.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "your message is invalid")
		return
	}

	dbMessage, err := h.postgres.EditMessage(c, c.Param("id"), message)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			NewErrorResponse(c, http.StatusNotFound, "Message does not exist")
			return
		default:
			logger.ErrorF("Error edit message: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

			NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	NewSuccessResponse(c, http.StatusOK, "Edit Message!", dbMessage)
}

func (h *Handlers) GetStats(c *gin.Context) {
	stats, err := h.postgres.GetStats(c)
	if err != nil {
		logger.InfoF("Error get statistic: %v", logrus.Fields{constants.LoggerCategory: constants.Handlers}, err)

		NewErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Here stats", stats)
}
