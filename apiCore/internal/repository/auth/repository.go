package auth

import (
	"database/sql"
	"go-mail-sender/apiCore/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	stmt, err := r.db.Prepare(createUserSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(findUserByIdSQL, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
