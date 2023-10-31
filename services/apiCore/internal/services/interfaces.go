package services

import (
	"context"
	"go-mail-sender/services/apiCore/internal/models"
	"mime/multipart"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterUser(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type FileService interface {
	CreateFile(file *multipart.FileHeader, userID uuid.UUID) (*models.FileResponse, error)
	GetFiles(ctx context.Context, useID uuid.UUID) ([]*models.File, error)
	GetFileByID(id, userID uuid.UUID) (*models.File, error)
}

type NewslettersService interface {
	CreateNewsletter(message string, userID uuid.UUID) (*models.Newsletter, error)
}
