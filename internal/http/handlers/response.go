package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(ctx *gin.Context, statusCode int, description string, data interface{}) {
	ctx.AbortWithStatusJSON(statusCode, Response{
		Status:      statusCode,
		Description: description,
		Data:        data,
	})
}

func NewErrorResponse(ctx *gin.Context, statusCode int, description string) {
	ctx.AbortWithStatusJSON(statusCode, Response{
		Status:      statusCode,
		Description: description,
	})
}
