package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Adiciona",
	})
}

func ListTaskHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Lista",
	})
}

func UpdateTaskHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Atualiza",
	})
}

func DeleteTaskHandler(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleta",
	})
}

func ListTasksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Listas",
	})
}
