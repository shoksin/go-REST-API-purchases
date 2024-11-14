package main

import (
	"github.com/joho/godotenv"
	"github.com/shoksin/go-REST-API-purchases/app"
	"log"
)

func main() {
	app.Run()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

}
