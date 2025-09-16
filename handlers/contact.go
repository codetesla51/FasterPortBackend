package handlers

import (
    "net/http"

    "github.com/codetesla51/portBackend/models"
    "github.com/gin-gonic/gin"
)

type MessageInput struct {
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required,email"`
    Inquiry string `json:"inquiry" binding:"required"`
    Message string `json:"message" binding:"required"`
}

func StoreMessageHandler(c *gin.Context) {
    var input models.MessageInput
if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

err := models.StoreMessage(input)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusCreated, gin.H{"message": "Message stored successfully"})
}