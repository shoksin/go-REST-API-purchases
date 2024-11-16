package db

import (
	"fmt"
	"github.com/shoksin/go-REST-API-purchases/config"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	config.Load()

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	fmt.Printf("host=%s user=%s dbname=%s sslmode=disable password=%s\n\n", dbHost, username, dbName, password)

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.GetLogger().Fatal("couldn't connect to database")
	}

	db = conn
	if err := db.Debug().AutoMigrate(&models.User{}, &models.Purchase{}); err != nil {
		logging.GetLogger().Fatal("unsuccessful database migration", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
