package config

import (
	"github.com/joho/godotenv"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
)

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		logging.GetLogger().Fatal("Error loading .env file")
	}
}
