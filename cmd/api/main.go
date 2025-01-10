package main

import (
	"github.com/shoksin/go-REST-API-purchases/app"
)

// @title Purchases API
// @version 1.0
// @description API Server for Purchases Tracker

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
