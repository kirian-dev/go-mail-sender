package main

import (
	"fmt"
	"go-mail-sender/config"
	"go-mail-sender/pkg/db"
	"go-mail-sender/pkg/log"
	subHttp "go-mail-sender/services/divider/internal/delivery/http"
	"go-mail-sender/services/divider/internal/repository"
	"go-mail-sender/services/divider/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	log := log.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error config:", err)
	}

	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHostDivider, cfg.PgPortDivider, cfg.PgUserDivider, cfg.PgPasswordDivider, cfg.PgDBNameDivider)

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

	sRepo := repository.NewSubscriberRepository(pgDB.DB)

	subServices := services.NewSubscribersService(sRepo)
	subHttp.SetupRoutes(apiV1, subServices, cfg, log)

	r.Run(":" + cfg.AppDividerPort)
}
