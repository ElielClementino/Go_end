package handler

import (
	"fmt"
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
	id := c.Query("id")
	if id == "" {
		SendError(c, http.StatusBadGateway, errParamIsRequired("id", "queryParam").Error())
		return
	}

	task := schemas.Task{}

	if err := db.First(&task, id).Error; err != nil {
		SendError(c, http.StatusNotFound, "Task not found")
		return
	}

	SendSuccess(c, "Task found successfully", task)
}

func UpdateTaskHandler(c *gin.Context) {
	request := UpdateTaskRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("error updating task %v", err.Error())
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")

	if id == "" {
		SendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	task := schemas.Task{}

	if err := db.First(&task, id).Error; err != nil {
		SendError(c, http.StatusBadRequest, "Task not found")
		return
	}

	if request.Title != "" {
		task.Title = request.Title
	}

	if request.Description != "" {
		task.Description = request.Description
	}

	if request.DueTo != "" {
		task.DueTo = request.DueTo
	}

	if err := db.Save(&task).Error; err != nil {
		logger.Errf("error while updating task: %v", err.Error())
		SendError(c, http.StatusInternalServerError, fmt.Sprintf("Error while updating task with id: %v", id))
		return
	}

	SendSuccess(c, "Task update successfully", task)
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		SendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		SendError(c, http.StatusNotFound, fmt.Sprintf("No object found with the respective id: %v", id))
		return
	}
	if err := db.Delete(&task).Error; err != nil {
		SendError(c, http.StatusInternalServerError, fmt.Sprintf("Error while deleting task with id: %v", id))
		return
	}
	SendSuccess(c, "Task with id: %v deleted successfully", id)
}

func ListTasksHandler(c *gin.Context) {
	tasks := []schemas.Task{}
	if err := db.Find(&tasks).Error; err != nil {
		logger.Errf("Error while finding tasks: %v", err.Error())
		SendError(c, http.StatusInternalServerError, fmt.Sprintln("Error while listing tasks"))
	}
	SendSuccess(c, "Listing all tasks", tasks)
}
