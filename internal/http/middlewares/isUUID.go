package middlewares

import (
	"Messaggio/internal/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func IsUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := uuid.Parse(id)
		if err != nil {
			handlers.NewErrorResponse(c, http.StatusBadRequest, "your id is not a valid uuid")
		}

		c.Next()
	}
}
