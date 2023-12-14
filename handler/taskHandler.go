package handler

import (
	"net/http"

	"github.com/ElielClementino/Go_end/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	request := CreateTaskRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	task := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		DueTo:       request.DueTo,
	}

	if err := db.Create(&task).Error; err != nil {
		logger.Errf("error creating task: %v", err.Error())
		SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendSuccess(c, "Task created successfully", task)
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
