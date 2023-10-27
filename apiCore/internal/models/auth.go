package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func NewUser(email, password string) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}
