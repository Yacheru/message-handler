package routes

import (
	"github.com/gin-gonic/gin"

	"Messaggio/internal/http/handlers"
)

type Route struct {
	handlers *handlers.Handlers
	router   *gin.RouterGroup
}

func NewRoute(router *gin.RouterGroup) *Route {
	return &Route{
		router: router,
	}
}

func (r *Route) Routes() {
	{
		r.router.GET("/", r.handlers.SayHello)
	}
}
