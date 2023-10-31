package newsletters

import (
	"encoding/json"
	"go-mail-sender/config"
	"go-mail-sender/pkg/jwt"
	"go-mail-sender/services/apiCore/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NewslettersHandler struct {
	newslettersService services.NewslettersService
	cfg                *config.Config
	log                *logrus.Logger
}

func NewNewslettersHandler(newslettersService services.NewslettersService, cfg *config.Config, log *logrus.Logger) *NewslettersHandler {
	return &NewslettersHandler{
		newslettersService: newslettersService,
		cfg:                cfg,
		log:                log,
	}
}

func (h *NewslettersHandler) CreateNewsletters(c *gin.Context) {
	userClaims, _ := c.Get("userClaims")
	userID := userClaims.(*jwt.CustomClaims).UserID

	type NewslettersReq struct {
		Message string `json:"message"`
	}

	var newsletterReq NewslettersReq
	err := json.NewDecoder(c.Request.Body).Decode(&newsletterReq)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newsletter, err := h.newslettersService.CreateNewsletter(newsletterReq.Message, userID)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newsletter)
}
