package main

import (
	"net/http"
	"os"
	"time"

	"github.com/codetesla51/portBackend/config"
	"github.com/codetesla51/portBackend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB() // initialize DB

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://devuthman.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Project routes
	router.GET("/projects", handlers.GetProjects)
	router.GET("/projects/:slug", handlers.GetProject)

	// Contact route
	router.POST("/contact", handlers.StoreMessageHandler)

	router.ServeHTTP(w, r)
}