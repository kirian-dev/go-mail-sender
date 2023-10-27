package services

import "go-mail-sender/apiCore/internal/models"

type AuthService interface {
	RegisterUser(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
}
