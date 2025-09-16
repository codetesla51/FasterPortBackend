// api/index.go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codetesla51/portBackend/config"
	"github.com/codetesla51/portBackend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	config.ConnectDB()

	router = gin.New()
	router.Use(gin.Recovery())

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://devuthman.vercel.app"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/projects", handlers.GetProjects)
	router.GET("/projects/:slug", handlers.GetProject)
	router.POST("/contact", handlers.StoreMessageHandler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}