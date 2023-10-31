package newsletters

import (
	"go-mail-sender/config"
	"go-mail-sender/services/divider/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.RouterGroup, newsletterService services.NewslettersServices, cfg *config.Config, log *logrus.Logger) {
	newsletterHandler := NewNewslettersHandler(newsletterService, cfg, log)

	newsletterGroup := r.Group("/newsletters")
	{
		newsletterGroup.POST("", newsletterHandler.CreateNewsletters)
	}
}
