package repository

import (
	"go-mail-sender/services/divider/internal/models"

	"github.com/google/uuid"
)

type SubscriberRepository interface {
	Create(subscriber *models.SubscriberRequest) error
	GetSubscriberCount() (int, error)
	FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error)
	GetSubscribersInBatches(batchSize int, userID uuid.UUID) ([]*models.Subscriber, error)
}

type NewslettersRepository interface {
	Create(message string, userID uuid.UUID) (*models.Newsletter, error)
}

type PacketsRepository interface {
	Create(subscribers []*models.Subscriber) error
}
