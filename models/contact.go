package models

import (
    "context"
    "errors"
"time"
    "github.com/codetesla51/portBackend/config"
)

type MessageInput struct {
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required,email"`
    Inquiry string `json:"inquiry" binding:"required"`
    Message string `json:"message" binding:"required"`
}

func StoreMessage(input MessageInput) error {
    query := `
    INSERT INTO contacts (name, email, inquiry, message, created_at)
    VALUES ($1, $2, $3, $4, $5)
`
_, err := config.DB.Exec(
    context.Background(),
    query,
    input.Name,
    input.Email,
    input.Inquiry,
    input.Message,
    time.Now(),
)
    if err != nil {
        return errors.New("message not sent: " + err.Error())
    }

    return nil
}