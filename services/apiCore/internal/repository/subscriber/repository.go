package subscriber

import (
	"database/sql"
	"go-mail-sender/services/apiCore/internal/models"
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

func (r *SubscriberRepository) Create(userID uuid.UUID, subscriber *models.Subscriber) error {
	subscriber.ID = uuid.New()
	subscriber.CreatedAt = time.Now().UTC()

	_, err := r.db.Exec(CreateSubscriber,
		subscriber.ID, subscriber.Email, subscriber.FirstName, subscriber.LastName, userID, subscriber.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *SubscriberRepository) FindByEmail(email string) (*models.Subscriber, error) {
	var subscriber models.Subscriber
	err := r.db.QueryRow(FindAccountByEmailSQL, email).Scan(&subscriber.ID, &subscriber.Email)
	if err != nil {
		return nil, err
	}
	return &subscriber, nil
}
