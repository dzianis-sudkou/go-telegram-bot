package app

import (
	"fmt"
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/client"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/database/postgres"
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

	// Database connection
	db := postgres.Connect()
	if db.AllowGlobalUpdate {
		fmt.Println("Connected!")
	}

	// Tables creation
	postgres.CreateTables(db)

	// Start the Bot
	client.Init()
}
