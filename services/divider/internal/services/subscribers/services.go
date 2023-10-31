package subscribers

import (
	"go-mail-sender/services/divider/internal/models"
	"go-mail-sender/services/divider/internal/repository"

	"github.com/google/uuid"
)

type SubscribersService struct {
	subscriberRepository repository.SubscriberRepository
}

func NewSubscribersService(subscriberRepository repository.SubscriberRepository) *SubscribersService {
	return &SubscribersService{
		subscriberRepository: subscriberRepository,
	}
}

func (s *SubscribersService) Create(subscriberReq *models.SubscriberRequest) error {
	err := s.subscriberRepository.Create(subscriberReq)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscribersService) FindByEmail(email string, userID uuid.UUID) (*models.Subscriber, error) {
	sub, err := s.subscriberRepository.FindByEmail(email, userID)
	if err != nil {
		return nil, err
	}
	return sub, nil
}
