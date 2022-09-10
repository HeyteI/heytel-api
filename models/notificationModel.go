package models

import "github.com/google/uuid"

type Notification struct {
	Base       `mapstructure:",squash"`
	Type       string    `json:"type" binding:"required"`    // Type of notification: "warning", "info", "alert"
	Message    string    `json:"message" binding:"required"` // Message of notification
	UserId     uuid.UUID `json:"user_id" binding:"required"` // ID of user that created notification
	ReceiverId uuid.UUID `json:"receiver_id"`                // ID of user that will receive notification (if empty, all users will receive)
}
