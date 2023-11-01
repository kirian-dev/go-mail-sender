package main

import (
	"fmt"
	"go-mail-sender/config"
	"go-mail-sender/pkg/db"
	"go-mail-sender/pkg/log"
	newHttp "go-mail-sender/services/divider/internal/delivery/http/newsletters"
	subHttp "go-mail-sender/services/divider/internal/delivery/http/subscribers"
	newRepo "go-mail-sender/services/divider/internal/repository/newsletters"
	packRepo "go-mail-sender/services/divider/internal/repository/packets"

	subRepo "go-mail-sender/services/divider/internal/repository/subscribers"
	newServ "go-mail-sender/services/divider/internal/services/newsletters"
	subServ "go-mail-sender/services/divider/internal/services/subscribers"

	"github.com/IBM/sarama"
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
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForLocal
	producerConfig.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		log.Fatal("Error initializing Kafka producer:", err)
	}
	defer producer.Close()
	r := gin.Default()
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	sRepo := subRepo.NewSubscriberRepository(pgDB.DB)
	newRepo := newRepo.NewNewslettersRepository(pgDB.DB)
	packRepo := packRepo.NewPacketsRepository(pgDB.DB)

	subServices := subServ.NewSubscribersService(sRepo)
	newServices := newServ.NewNewslettersService(newRepo, packRepo, sRepo, producer, cfg, log)

	subHttp.SetupRoutes(apiV1, subServices, cfg, log)
	newHttp.SetupRoutes(apiV1, newServices, cfg, log)

	r.Run(":" + cfg.AppDividerPort)
}
