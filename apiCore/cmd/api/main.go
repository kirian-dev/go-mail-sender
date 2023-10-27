package main

import (
	"fmt"
	authHttp "go-mail-sender/apiCore/internal/delivery/http/auth"
	authServices "go-mail-sender/apiCore/internal/services/auth"

	authRepo "go-mail-sender/apiCore/internal/repository/auth"

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

	aRepo := authRepo.NewAuthRepository(pgDB.DB)
	aServices := authServices.NewAuthService(aRepo)

	authHttp.SetupRoutes(apiV1, aServices, cfg, log)

	r.Run(":" + cfg.AppPort)
}
