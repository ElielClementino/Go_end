package router

import (
	"github.com/ElielClementino/Go_end/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("add/task", handler.CreateTaskHandler)
		api.GET("list/task", handler.ListTaskHandler)
		api.PUT("update/task", handler.UpdateTaskHandler)
		api.DELETE("delete/task", handler.DeleteTaskHandler)
		api.GET("list/tasks", handler.ListTasksHandler)
	}
}
