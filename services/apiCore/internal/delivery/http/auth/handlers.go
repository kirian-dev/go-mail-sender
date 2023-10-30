package auth

import (
	"go-mail-sender/config"
	"go-mail-sender/pkg/jwt"
	"go-mail-sender/services/apiCore/internal/models"
	"go-mail-sender/services/apiCore/internal/services"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService services.AuthService
	cfg         *config.Config
	log         *logrus.Logger
}

func NewAuthHandler(authService services.AuthService, cfg *config.Config, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		cfg:         cfg,
		log:         log,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if len(user.Password) < 6 || len(user.Password) > 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password length must be between 6 and 32 characters"})
		return
	}

	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := jwt.GenerateAccessToken(createdUser.Email, createdUser.ID, h.cfg)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := checkmail.ValidateFormat(LoginRequest.Email); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if len(LoginRequest.Password) < 6 || len(LoginRequest.Password) > 32 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Password length must be between 6 and 32 characters"})
		return
	}
	user, err := h.authService.Login(LoginRequest.Email, LoginRequest.Password)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := jwt.GenerateAccessToken(user.Email, user.ID, h.cfg)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
