package file

import (
	"go-mail-sender/config"
	"go-mail-sender/services/apiCore/internal/middleware"
	"go-mail-sender/services/apiCore/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(r *gin.RouterGroup, fileService services.FileService, cfg *config.Config, log *logrus.Logger, mw *middleware.MiddlewareManager) {
	fileHandler := NewFileHandler(fileService, cfg, log)

	fileGroup := r.Group("/files")
	{
		fileGroup.POST("", mw.AuthMiddleware(), fileHandler.CreateFile)
		fileGroup.GET("", mw.AuthMiddleware(), fileHandler.GetFiles)
		fileGroup.GET("/:fileId", mw.AuthMiddleware(), fileHandler.GetFileByID)
	}
}
