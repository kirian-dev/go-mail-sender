package main

import (
	"fmt"
	authHttp "go-mail-sender/services/apiCore/internal/delivery/http/auth"
	fileHttp "go-mail-sender/services/apiCore/internal/delivery/http/file"

	"go-mail-sender/services/apiCore/internal/middleware"
	authServices "go-mail-sender/services/apiCore/internal/services/auth"
	fileServices "go-mail-sender/services/apiCore/internal/services/file"

	authRepo "go-mail-sender/services/apiCore/internal/repository/auth"
	fileRepo "go-mail-sender/services/apiCore/internal/repository/file"

	"go-mail-sender/config"
	"go-mail-sender/pkg/db"
	"go-mail-sender/pkg/log"

	"github.com/gin-gonic/gin"
)

func main() {
	log := log.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error config:", err)
	}

	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDBName)

	pgDB, err := db.Connect("postgres", pgConnStr)
	if err != nil {
		log.Fatal("Error connect to db", err)
	}
	defer pgDB.Close()

	if err = pgDB.Ping(); err != nil {
		log.Fatal("Error check connect to db:", err)
	}

	log.Info("DB connection is successful")

	r := gin.Default()
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mw := middleware.NewMiddlewareManager(cfg, log)

	aRepo := authRepo.NewAuthRepository(pgDB.DB)
	fRepo := fileRepo.NewFileRepository(pgDB.DB)

	aServices := authServices.NewAuthService(aRepo)
	fServices := fileServices.NewFileService(fRepo, log, cfg)
	authHttp.SetupRoutes(apiV1, aServices, cfg, log)

	fileHttp.SetupRoutes(apiV1, fServices, cfg, log, mw)

	r.Run(":" + cfg.AppPort)
}
