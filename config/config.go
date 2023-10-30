package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	PgHost         string
	PgPort         string
	PgUser         string
	PgPassword     string
	PgDBName       string
	JWTKey         string
	UploadFolder   string
	BufferSize     string
	GoroutineCount string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cfg := &Config{
		AppPort:        os.Getenv("APP_PORT"),
		PgHost:         os.Getenv("PG_HOST"),
		PgPort:         os.Getenv("PG_PORT"),
		PgUser:         os.Getenv("PG_USER"),
		PgPassword:     os.Getenv("PG_PASSWORD"),
		PgDBName:       os.Getenv("PG_DBNAME"),
		JWTKey:         os.Getenv("JWT_KEY"),
		UploadFolder:   os.Getenv("UPLOAD_FOLDER"),
		BufferSize:     os.Getenv("BUFFER_SIZE"),
		GoroutineCount: os.Getenv("GOROUTINE_COUNT"),
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	requiredFields := []struct {
		value interface{}
		name  string
	}{
		{cfg.PgHost, "APP_PORT"},
		{cfg.PgHost, "PG_HOST"},
		{cfg.PgPort, "PG_PORT"},
		{cfg.PgUser, "PG_USER"},
		{cfg.PgPassword, "PG_PASSWORD"},
		{cfg.PgDBName, "PG_DBNAME"},
		{cfg.JWTKey, "JWT_KEY"},
		{cfg.UploadFolder, "UPLOAD_FOLDER"},
		{cfg.BufferSize, "BUFFER_SIZE"},
		{cfg.GoroutineCount, "GOROUTINE_COUNT"},
	}

	for _, field := range requiredFields {
		if field.value == "" {
			return fmt.Errorf("configuration variable '%s' must not be missing or empty", field.name)
		}
	}

	return nil
}
