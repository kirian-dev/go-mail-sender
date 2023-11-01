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
	UpdatePacketID(subscriberID, packetID uuid.UUID) error
	UpdatePacketIDForSubscribers(subscribers []*models.Subscriber, packetID uuid.UUID) error
}

type NewslettersRepository interface {
	Create(message string, userID uuid.UUID) (*models.Newsletter, error)
}

type PacketsRepository interface {
	Create(msg string) (*models.Packet, error)
}
