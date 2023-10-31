package services

import (
	"go-mail-sender/services/divider/internal/models"

	"github.com/google/uuid"
)

type ISubscriberServices interface {
	Create(subscriberReq *models.SubscriberRequest) error
	FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error)
}
