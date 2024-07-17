package routes

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Handlers []gin.HandlerFunc
	router   *gin.RouterGroup
}

func (r *Route) NewRoute() *Route {
	return &Route{}
}

func (r *Route) Routes() {
	{
		r.router.POST("/")
		r.router.GET("/")
	}
}
