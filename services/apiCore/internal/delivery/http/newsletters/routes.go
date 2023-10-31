package newsletters

import (
	"go-mail-sender/config"
	"go-mail-sender/services/apiCore/internal/middleware"
	"go-mail-sender/services/apiCore/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.RouterGroup, newsletterService services.NewslettersService, cfg *config.Config, log *logrus.Logger, mw *middleware.MiddlewareManager) {
	newsletterHandler := NewNewslettersHandler(newsletterService, cfg, log)

	newsletterGroup := r.Group("/newsletters")
	{
		newsletterGroup.POST("", mw.AuthMiddleware(), newsletterHandler.CreateNewsletters)
	}
}
