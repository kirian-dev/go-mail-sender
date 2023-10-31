package http

import (
	"go-mail-sender/config"
	"go-mail-sender/services/divider/internal/models"
	"go-mail-sender/services/divider/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SubscriberHandler struct {
	subscriberService services.ISubscriberServices
	cfg               *config.Config
	log               *logrus.Logger
}

func NewSubscriberHandler(subscriberService services.ISubscriberServices, cfg *config.Config,
	log *logrus.Logger) *SubscriberHandler {
	return &SubscriberHandler{subscriberService: subscriberService, cfg: cfg, log: log}
}

func (h *SubscriberHandler) CreateSubscriber(c *gin.Context) {
	var subscriberReq models.SubscriberRequest
	if err := c.ShouldBindJSON(&subscriberReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.subscriberService.Create(&subscriberReq)
	h.log.Info(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subscriberReq)
}

func (h *SubscriberHandler) GetSubscriber(c *gin.Context) {
	subscriberEmail := c.Param("subscriberEmail")
	userID := c.Query("user_id")
	parsedID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriber, err := h.subscriberService.FindByEmail(subscriberEmail, parsedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriber)
}
