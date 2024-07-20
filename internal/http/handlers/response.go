package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, description string, data interface{}) {
	c.AbortWithStatusJSON(statusCode, Response{
		Status:      statusCode,
		Description: description,
		Data:        data,
	})
}

func NewErrorResponse(c *gin.Context, statusCode int, description string) {
	c.AbortWithStatusJSON(statusCode, Response{
		Status:      statusCode,
		Description: description,
	})
}
