package repository

import "go-mail-sender/apiCore/internal/models"

type AuthRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}
