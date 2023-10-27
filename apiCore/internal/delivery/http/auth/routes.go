package auth

import (
	"go-mail-sender/apiCore/internal/services"
	"go-mail-sender/config"

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
