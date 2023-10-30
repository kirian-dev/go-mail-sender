package auth

import (
	"go-mail-sender/config"
	"go-mail-sender/services/apiCore/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.RouterGroup, authService services.AuthService, cfg *config.Config, log *logrus.Logger) {
	authHandler := NewAuthHandler(authService, cfg, log)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.RegisterUser)
		authGroup.POST("/login", authHandler.Login)
	}
}
