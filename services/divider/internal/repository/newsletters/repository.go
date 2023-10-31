package newsletters

import (
	"database/sql"
	"go-mail-sender/services/divider/internal/models"
	"time"

	"github.com/google/uuid"
)

type NewslettersRepository struct {
	db *sql.DB
}

func NewNewslettersRepository(db *sql.DB) *NewslettersRepository {
	return &NewslettersRepository{
		db: db,
	}
}

func (r *NewslettersRepository) Create(message string, userID uuid.UUID) (*models.Newsletter, error) {
	newsletterID := uuid.New()
	newsletter := &models.Newsletter{
		ID:        newsletterID,
		Message:   message,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
	}

	_, err := r.db.Exec(CreateNewsletter, newsletterID, message, userID, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return newsletter, nil
}
