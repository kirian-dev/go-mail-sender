package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(email, password string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}
