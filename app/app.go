package app

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/config"
	db2 "github.com/shoksin/go-REST-API-purchases/internal/db"
	"github.com/shoksin/go-REST-API-purchases/internal/handlers"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"os"
)

const (
	PORT = "PORT"
	HOST = "HOST"
)

func Run() {
	config.Load()
	address := os.Getenv(HOST) + ":" + os.Getenv(PORT)

	e := echo.New()

	db := db2.GetDB()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	e.POST("/register", userHandler.CreateUser)

	if err := e.Start(address); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
