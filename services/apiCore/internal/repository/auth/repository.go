package auth

import (
	"database/sql"
	"go-mail-sender/services/apiCore/internal/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user *models.User) (*models.User, error) {
	stmt, err := r.db.Prepare(createUserSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password, user.ID, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(findUserByEmailSQL, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
