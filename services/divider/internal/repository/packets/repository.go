package packets

import (
	"database/sql"
	"go-mail-sender/services/divider/internal/models"
	"time"

	"github.com/google/uuid"
)

type PacketsRepository struct {
	db *sql.DB
}

func NewPacketsRepository(db *sql.DB) *PacketsRepository {
	return &PacketsRepository{
		db: db,
	}
}

func (r *PacketsRepository) Create(subscribers []*models.Subscriber) error {
	_, err := r.db.Exec(CreatePacket, uuid.New(), time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}
