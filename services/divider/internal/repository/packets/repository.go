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

func (r *PacketsRepository) Create(msg string) (*models.Packet, error) {
	packet := &models.Packet{
		ID:        uuid.New(),
		Message:   msg,
		CreatedAt: time.Now().UTC(),
	}

	_, err := r.db.Exec(CreatePacket, packet.ID, packet.Message, packet.CreatedAt)

	if err != nil {
		return nil, err
	}

	return packet, nil
}
