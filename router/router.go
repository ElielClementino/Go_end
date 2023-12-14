package router

import (
	"github.com/ElielClementino/Go_end/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	// Initialize handler
	handler.InitializeHandler()

	// Initialize router with default confings
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Initialize routes
	InitializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
