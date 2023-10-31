package NewsLetters

import (
	"go-mail-sender/config"
	"go-mail-sender/services/divider/internal/models"
	"go-mail-sender/services/divider/internal/repository"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const workerCount = 4

type NewslettersService struct {
	newslettersRepository repository.NewslettersRepository
	packetsRepository     repository.PacketsRepository
	subscribersRepository repository.SubscriberRepository
	cfg                   *config.Config
	log                   *logrus.Logger
}

func NewNewslettersService(newslettersRepository repository.NewslettersRepository, packetsRepository repository.PacketsRepository, cfg *config.Config, log *logrus.Logger) *NewslettersService {
	return &NewslettersService{
		cfg:                   cfg,
		newslettersRepository: newslettersRepository,
		packetsRepository:     packetsRepository,
		log:                   log,
	}
}

func (s *NewslettersService) Create(message string, userID uuid.UUID) (*models.Newsletter, error) {
	subscribers, err := s.subscribersRepository.GetSubscribersInBatches(1000, userID)
	if err != nil {
		return nil, err
	}

	newsletter, err := s.newslettersRepository.Create(message, userID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	tasks := make(chan []*models.Subscriber, workerCount)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range tasks {
				s.createWorker(batch)
			}
		}()
	}

	for i := 0; i < len(subscribers); i += workerCount {
		end := i + workerCount
		if end > len(subscribers) {
			end = len(subscribers)
		}
		tasks <- subscribers[i:end]
	}

	close(tasks)

	wg.Wait()

	return newsletter, nil
}

func (s *NewslettersService) createWorker(subscribers []*models.Subscriber) {
	for _, subscriber := range subscribers {

		s.log.Infof("Sending email to: %s", subscriber.Email)
	}
}
