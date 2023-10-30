package repository

import (
	"context"
	"go-mail-sender/services/apiCore/internal/models"

	"github.com/google/uuid"
)

type AuthRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

type FileRepository interface {
	CreateFile(fileName string, userID uuid.UUID, rows int) (*models.File, error)
	UpdateFile(file *models.File) (*models.File, error)
	FindFileByID(id, userID uuid.UUID) (*models.File, error)
	GetFiles(ctx context.Context, userID uuid.UUID) ([]*models.File, error)
}

type SubscriberRepository interface {
	Create(userID uuid.UUID, subscriber *models.Subscriber) error
	GetSubscriberCount() (int, error)
	FindByEmail(email string) (*models.Subscriber, error)
}
