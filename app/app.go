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

	db := db2.GetDB()
	userRepository := repositories.NewUserRepository(db)
	purchasesRepository := repositories.NewPurchasesRepository(db)
	userService := services.NewUserService(userRepository)
	purchasesService := services.NewPurchasesService(purchasesRepository)
	userHandler := handlers.NewUserHandler(userService)
	purchasesHandler := handlers.NewPurchasesHandler(purchasesService)

	e.POST("/register", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)
	e.POST("/purchases/add", purchasesHandler.CreatePurchase)
	e.GET("/purchases/get", purchasesHandler.GetPurchases)
	e.DELETE("/purchases/delete", purchasesHandler.DeletePurchase)
	e.DELETE("/purchases/delete", purchasesHandler.DeleteUserPurchases)

	if err := e.Start(address); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
