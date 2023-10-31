package subscribers

import (
	"go-mail-sender/config"
	"go-mail-sender/services/divider/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.RouterGroup, subscribersService services.SubscriberServices, cfg *config.Config, log *logrus.Logger) {
	subscribersHandler := NewSubscriberHandler(subscribersService, cfg, log)

	subscribersGroup := r.Group("/subscribers")
	{
		subscribersGroup.POST("", subscribersHandler.CreateSubscriber)
		subscribersGroup.GET("/:subscriberEmail", subscribersHandler.GetSubscriber)
	}
}
