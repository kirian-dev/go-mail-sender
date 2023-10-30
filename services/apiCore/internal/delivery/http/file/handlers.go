package file

import (
	"go-mail-sender/config"
	"go-mail-sender/pkg/jwt"
	"go-mail-sender/services/apiCore/internal/services"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	maxFileSize = 2024 * 2024
	allowedExt  = ".csv"
)

type FileHandler struct {
	fileService services.FileService
	cfg         *config.Config
	log         *logrus.Logger
}

func NewFileHandler(fileService services.FileService, cfg *config.Config, log *logrus.Logger) *FileHandler {
	return &FileHandler{
		fileService: fileService,
		cfg:         cfg,
		log:         log,
	}
}

func (h *FileHandler) GetFiles(c *gin.Context) {
	userClaims, _ := c.Get("userClaims")
	userID := userClaims.(*jwt.CustomClaims).UserID

	exists, err := h.fileService.GetFiles(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"file": exists})
}

func (h *FileHandler) GetFileByID(c *gin.Context) {
	fileStrID := c.Param(":fileId")
	fileID, err := uuid.Parse(fileStrID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	userClaims, _ := c.Get("userClaims")
	userID := userClaims.(*jwt.CustomClaims).UserID

	exists, err := h.fileService.GetFileByID(fileID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": exists})
}

func (h *FileHandler) CreateFile(c *gin.Context) {
	// Retrieve user email from the JWT token
	userClaims, _ := c.Get("userClaims")
	userID := userClaims.(*jwt.CustomClaims).UserID

	// File validation
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	// Perform file size and extension validation
	if file.Size > maxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File size is too large"})
		return
	}
	ext := filepath.Ext(file.Filename)
	if ext != allowedExt {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
		return
	}
	createdFile, err := h.fileService.CreateFile(file, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	successResponse := struct {
		Message string    `json:"message"`
		ID      uuid.UUID `json:"id"`
		Name    string    `json:"name"`
	}{
		Message: "File uploaded successfully",
		ID:      createdFile.ID,
		Name:    createdFile.Name,
	}

	c.JSON(http.StatusCreated, successResponse)
}
