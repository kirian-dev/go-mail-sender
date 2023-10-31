package services

import (
	"go-mail-sender/services/divider/internal/models"

	"github.com/google/uuid"
)

type SubscriberServices interface {
	Create(subscriberReq *models.SubscriberRequest) error
	FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error)
}

type NewslettersServices interface {
	Create(message string, userID uuid.UUID) (*models.Newsletter, error)
}
