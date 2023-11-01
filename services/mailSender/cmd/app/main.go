package main

import (
	"go-mail-sender/config"
	"go-mail-sender/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log := log.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error config:", err)
	}

	r := gin.Default()
	apiV1 := r.Group("/api/v1")

	apiV1.POST("/send-email", func(c *gin.Context) {
		var request struct {
			Email   string `json:"email"`
			Message string `json:"message"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		log.Infof("Sending email to: %s, Message: %s", request.Email, request.Message)

		c.JSON(http.StatusOK, gin.H{
			"message": "Email sent",
		})
	})

	r.Run(":" + cfg.AppMailSenderPort)
	log.Info("Microservice mail sender started")
}
