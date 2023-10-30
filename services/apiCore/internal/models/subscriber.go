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
