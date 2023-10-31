package models

import (
	"time"

	"github.com/google/uuid"
)

type Newsletter struct {
	ID        uuid.UUID
	Message   string
	UserID    uuid.UUID
	CreatedAt time.Time
}

type NewsletterRequest struct {
	Message string    `json:"message"`
	UserID  uuid.UUID `json:"user_id"`
}
