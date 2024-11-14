package app

import (
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/config"
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
	if err := e.Start(address); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
