package utils

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func SendErrorResponse(context *gin.Context, statusCode int, status string, message string, errors interface{}) {
	context.JSON(statusCode, ErrorResponse{
		Status:  status,
		Message: message,
		Error:   errors,
	})
}

func SendSuccessResponse(context *gin.Context, statusCode int, status string, message string, data interface{}) {
	context.JSON(statusCode, SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
