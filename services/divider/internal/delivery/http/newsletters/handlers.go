package newsletters

import (
	"go-mail-sender/config"
	"go-mail-sender/services/divider/internal/models"
	"go-mail-sender/services/divider/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NewslettersHandler struct {
	newslettersService services.NewslettersServices
	cfg                *config.Config
	log                *logrus.Logger
}

func NewNewslettersHandler(newslettersService services.NewslettersServices, cfg *config.Config, log *logrus.Logger) *NewslettersHandler {
	return &NewslettersHandler{
		newslettersService: newslettersService,
		cfg:                cfg,
		log:                log,
	}
}

func (h *NewslettersHandler) CreateNewsletters(c *gin.Context) {
	var newsletterReq models.NewsletterRequest
	if err := c.ShouldBindJSON(&newsletterReq); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.log.Info(&newsletterReq)
	newsletter, err := h.newslettersService.Create(newsletterReq.Message, newsletterReq.UserID)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, newsletter)
}
