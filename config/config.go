package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort           string
	PgHost            string
	PgPort            string
	PgUser            string
	PgPassword        string
	PgDBName          string
	JWTKey            string
	UploadFolder      string
	BufferSize        string
	GoroutineCount    string
	AppDividerPort    string
	AppDividerURL     string
	PgHostDivider     string
	PgPortDivider     string
	PgUserDivider     string
	PgPasswordDivider string
	PgDBNameDivider   string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cfg := &Config{
		AppPort:           os.Getenv("APP_API_CORE_PORT"),
		PgHost:            os.Getenv("PG_HOST"),
		PgPort:            os.Getenv("PG_PORT"),
		PgUser:            os.Getenv("PG_USER"),
		PgPassword:        os.Getenv("PG_PASSWORD"),
		PgDBName:          os.Getenv("PG_DBNAME"),
		JWTKey:            os.Getenv("JWT_KEY"),
		UploadFolder:      os.Getenv("UPLOAD_FOLDER"),
		BufferSize:        os.Getenv("BUFFER_SIZE"),
		GoroutineCount:    os.Getenv("GOROUTINE_COUNT"),
		AppDividerPort:    os.Getenv("APP_DIVIDER_PORT"),
		AppDividerURL:     os.Getenv("APP_DIVIDER_URL"),
		PgHostDivider:     os.Getenv("PG_HOST_DIVIDER"),
		PgPortDivider:     os.Getenv("PG_PORT_DIVIDER"),
		PgUserDivider:     os.Getenv("PG_USER_DIVIDER"),
		PgPasswordDivider: os.Getenv("PG_PASSWORD_DIVIDER"),
		PgDBNameDivider:   os.Getenv("PG_DBNAME_DIVIDER"),
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
		{cfg.AppPort, "APP_API_CORE_PORT"},
		{cfg.PgHost, "PG_HOST"},
		{cfg.PgPort, "PG_PORT"},
		{cfg.PgUser, "PG_USER"},
		{cfg.PgPassword, "PG_PASSWORD"},
		{cfg.PgDBName, "PG_DBNAME"},
		{cfg.JWTKey, "JWT_KEY"},
		{cfg.UploadFolder, "UPLOAD_FOLDER"},
		{cfg.BufferSize, "BUFFER_SIZE"},
		{cfg.GoroutineCount, "GOROUTINE_COUNT"},
		{cfg.AppDividerPort, "APP_DIVIDER_PORT"},
		{cfg.AppDividerURL, "APP_DIVIDER_URL"},
		{cfg.PgHostDivider, "PG_HOST_DIVIDER"},
		{cfg.PgPortDivider, "PG_POR_DIVIDERT"},
		{cfg.PgUserDivider, "PG_USER_DIVIDER"},
		{cfg.PgPasswordDivider, "PG_PASSWORD_DIVIDER"},
		{cfg.PgDBNameDivider, "PG_DBNAME_DIVIDER"},
	}

	for _, field := range requiredFields {
		if field.value == "" {
			return fmt.Errorf("configuration variable '%s' must not be missing or empty", field.name)
		}
	}

	return nil
}
