package main

import (
	"os"
	"time"

	"github.com/codetesla51/portBackend/config"
	"github.com/codetesla51/portBackend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://your-frontend-domain.com", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Project routes
	r.GET("/projects", handlers.GetProjects)
	r.GET("/projects/:slug", handlers.GetProject)

	// Contact route
	r.POST("/contact", handlers.StoreMessageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}