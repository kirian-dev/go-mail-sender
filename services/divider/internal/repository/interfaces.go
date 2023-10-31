package repository

import (
	"go-mail-sender/services/divider/internal/models"

	"github.com/google/uuid"
)

type ISubscriberRepository interface {
	Create(subscriber *models.SubscriberRequest) error
	GetSubscriberCount() (int, error)
	FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error)
}
