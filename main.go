package main

import (
	"github.com/codetesla51/portBackend/config"
	"github.com/codetesla51/portBackend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/projects", handlers.GetProjects)
	r.GET("/projects/:slug", handlers.GetProject)
	r.POST("/contact", handlers.StoreMessageHandler)

	r.Run(":8080")
}
