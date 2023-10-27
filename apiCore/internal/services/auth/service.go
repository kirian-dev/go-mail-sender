package auth

import (
	"errors"
	"go-mail-sender/apiCore/internal/models"
	"go-mail-sender/apiCore/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	invalidCredentials = errors.New("invalid credentials")
	emailExists        = errors.New("email exists")
)

type AuthService struct {
	userRepository repository.AuthRepository
}

func NewAuthService(userRepository repository.AuthRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) RegisterUser(user *models.User) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	existsUser, _ := s.userRepository.FindUserByEmail(user.Email)

	if existsUser != nil {
		return nil, emailExists
	}
	user = models.NewUser(user.Email, string(hashedPassword))

	createdUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, invalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, invalidCredentials
	}

	return user, nil
}
