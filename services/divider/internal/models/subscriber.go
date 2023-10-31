package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscriber struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	UserID    uuid.UUID
	CreatedAt time.Time
}

type SubscriberRequest struct {
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserID    uuid.UUID `json:"user_id"`
}
