package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse adalah struktur umum untuk respons JSON
type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse mengirimkan respons sukses dengan status HTTP 200
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, JSONResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// CreatedResponse mengirimkan respons sukses untuk status HTTP 201
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, JSONResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ErrorResponse mengirimkan respons error dengan status HTTP yang sesuai
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, JSONResponse{
		Status:  "error",
		Message: message,
	})
}

// BadRequestResponse mengirimkan respons error untuk status HTTP 400
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

// InternalServerErrorResponse mengirimkan respons error untuk status HTTP 500
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}
