package models

import (
	"time"

	"github.com/google/uuid"
)

type Packet struct {
	ID          uuid.UUID
	Subscribers []Subscriber
	CreatedAt   time.Time
}
