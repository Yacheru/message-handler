package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
}

func (h *Handlers) SayHello(c *gin.Context) {
	NewSuccessResponse(c, http.StatusOK, "Hello!", struct{}{})
}
