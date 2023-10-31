package NewsLetters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mail-sender/config"
	"go-mail-sender/services/apiCore/internal/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NewslettersService struct {
	cfg *config.Config
	log *logrus.Logger
}

func NewNewslettersService(cfg *config.Config, log *logrus.Logger) *NewslettersService {
	return &NewslettersService{
		cfg: cfg,
		log: log,
	}
}

func (s *NewslettersService) CreateNewsletter(message string, userID uuid.UUID) (*models.Newsletter, error) {
	type NewsletterBody struct {
		Message string    `json:"message"`
		UserID  uuid.UUID `json:"user_id"`
	}

	newsletterBody := &NewsletterBody{
		Message: message,
		UserID:  userID,
	}

	data, err := json.Marshal(newsletterBody)
	if err != nil {
		s.log.Errorf("Failed to marshal newsletter data: %v", err)
		return nil, err
	}
	url := fmt.Sprintf("%s:%s/api/v1/newsletters", s.cfg.AppDividerURL, s.cfg.AppDividerPort)
	client := &http.Client{}
	s.log.Info(string(data))
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		s.log.Errorf("Failed to send POST request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		s.log.Errorf("Failed to create newsletter in divider. Status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to create newsletter in divider. Status code: %d", resp.StatusCode)
	}

	var newsletter models.Newsletter
	err = json.NewDecoder(resp.Body).Decode(&newsletter)
	if err != nil {
		return nil, err
	}

	s.log.Infof("Newsletter created in divider: %s", newsletter.ID)
	return &newsletter, nil
}
