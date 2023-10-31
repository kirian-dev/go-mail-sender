package repository

import (
	"database/sql"
	"errors"
	"go-mail-sender/services/divider/internal/models"
	"time"

	"github.com/google/uuid"
)

type SubscriberRepository struct {
	db *sql.DB
}

func NewSubscriberRepository(db *sql.DB) *SubscriberRepository {
	return &SubscriberRepository{
		db: db,
	}
}

func (r *SubscriberRepository) GetSubscriberCount() (int, error) {
	var count int
	err := r.db.QueryRow(GetSubscriberCount).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *SubscriberRepository) Create(subscriberReq *models.SubscriberRequest) error {
	_, err := r.db.Exec(CreateSubscriber, uuid.New(), subscriberReq.Email, subscriberReq.FirstName, subscriberReq.LastName, subscriberReq.UserID, time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *SubscriberRepository) FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error) {
	var subscriber models.Subscriber
	err := r.db.QueryRow(FindAccountByEmailSQL, email, userID).Scan(&subscriber.ID, &subscriber.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("subscriber does not exists")
		}
		return nil, err
	}
	return &subscriber, nil
}
