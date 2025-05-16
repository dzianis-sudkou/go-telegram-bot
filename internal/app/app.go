package app

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/client"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/database/postgres"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	"github.com/joho/godotenv"
)

func Run() {

	/*
		SETUP
	*/

	// Import env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Read .env file error: %v", err)
	}

	// Start the DB connection
	repositories.DB = postgres.Init()

	// Start the Bot
	client.Init()
}
