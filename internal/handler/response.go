package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, baseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func newErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, baseResponse{
		Status:  false,
		Message: err,
	})
}

func newAbortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, baseResponse{
		Status:  false,
		Message: message,
	})
}
