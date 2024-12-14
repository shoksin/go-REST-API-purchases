package app

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/config"
	db2 "github.com/shoksin/go-REST-API-purchases/internal/db"
	"github.com/shoksin/go-REST-API-purchases/internal/handlers"
	"github.com/shoksin/go-REST-API-purchases/internal/repositories"
	"github.com/shoksin/go-REST-API-purchases/internal/services"
	"github.com/shoksin/go-REST-API-purchases/middleware"
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

	e.Use(middleware.JWTAuth)

	logger := logging.GetLogger()
	db := db2.GetDB()
	userRepository := repositories.NewUserRepository(db, logger)
	purchasesRepository := repositories.NewPurchasesRepository(db, logger)
	userService := services.NewUserService(userRepository, logger)
	purchasesService := services.NewPurchasesService(purchasesRepository, logger)
	userHandler := handlers.NewUserHandler(userService, logger)
	purchasesHandler := handlers.NewPurchasesHandler(purchasesService, logger)

	e.POST("/register", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)
	e.POST("/purchases/", purchasesHandler.CreatePurchase)
	e.GET("/purchases/", purchasesHandler.GetPurchases)
	e.DELETE("/purchases/", purchasesHandler.DeletePurchase)
	e.DELETE("/purchases/", purchasesHandler.DeleteUserPurchases)

	if err := e.Start(address); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
