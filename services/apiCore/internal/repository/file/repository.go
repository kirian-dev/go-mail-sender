package file

import (
	"context"
	"database/sql"
	"go-mail-sender/services/apiCore/internal/models"
	"time"

	"github.com/google/uuid"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}
func (r *FileRepository) GetFiles(ctx context.Context, userID uuid.UUID) ([]*models.File, error) {
	var files []*models.File
	rows, err := r.db.QueryContext(ctx, GetFilesSQL, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.ID, &file.UserID, &file.Name, &file.SuccessAccounts, &file.FailAccounts, &file.LoadingAccounts, &file.CreatedAt, &file.EndTime, &file.Status)
		if err != nil {
			return nil, err
		}
		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *FileRepository) FindFileByID(ID, useID uuid.UUID) (*models.File, error) {
	var file models.File
	err := r.db.QueryRow(FindFileByIDSQL, ID, useID).Scan(&file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) CreateFile(fileName string, userID uuid.UUID, rows int) (*models.File, error) {
	var newFile models.File

	err := r.db.QueryRow(CreateFileSQL, uuid.New(), userID, fileName, 0, 0, rows, time.Now(), time.Time{}, models.LoadingStatus).Scan(&newFile.ID)
	if err != nil {
		return nil, err
	}

	newFile.UserID = userID
	newFile.Name = fileName
	newFile.SuccessAccounts = 0
	newFile.FailAccounts = 0
	newFile.LoadingAccounts = rows
	newFile.CreatedAt = time.Now()
	newFile.EndTime = time.Time{}
	newFile.Status = models.LoadingStatus

	return &newFile, nil
}

func (r *FileRepository) UpdateFile(file *models.File) (*models.File, error) {
	_, err := r.db.Exec(UpdateFileSQL, file.ID, file.SuccessAccounts, file.FailAccounts, file.LoadingAccounts, file.EndTime, file.Status)
	if err != nil {
		return nil, err
	}

	return file, nil
}
