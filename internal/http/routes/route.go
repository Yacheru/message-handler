package routes

import (
	"Messaggio/internal/kafka/consumer"
	"Messaggio/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"Messaggio/internal/http/handlers"
	"Messaggio/internal/http/middlewares"
	"Messaggio/internal/kafka/producer"
	"Messaggio/internal/repository"
)

type Route struct {
	handlers *handlers.Handlers
	router   *gin.RouterGroup
}

func NewRoute(ctx context.Context, router *gin.RouterGroup, producer *producer.Producer, db *sqlx.DB, topic []string) *Route {
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandlers(producer, services)
	_ = consumer.NewConsumerGroup(ctx, topic, services)

	return &Route{
		handlers: handler,
		router:   router,
	}
}

func (r *Route) Routes() {
	{
		r.router.POST("/", r.handlers.InsertNew)                                // Добавить сообщение
		r.router.GET("/", r.handlers.GetAll)                                    // Получить все сообщения
		r.router.GET("/stats", r.handlers.GetStats)                             // Получить статистику
		r.router.GET("/:id", middlewares.IsUUID(), r.handlers.GetByID)          // Получить сообщение по его ID
		r.router.DELETE("/:id", middlewares.IsUUID(), r.handlers.DeleteMessage) // Удалить сообщение по ID
		r.router.PATCH("/:id", middlewares.IsUUID(), r.handlers.EditMessage)    // Изменить сообщение по ID
	}
}
