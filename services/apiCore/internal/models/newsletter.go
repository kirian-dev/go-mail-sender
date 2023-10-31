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
