package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int, message string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
	})
}

func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
