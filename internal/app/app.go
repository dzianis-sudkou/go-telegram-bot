package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/client"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/database/postgres"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/websocket"
	"github.com/joho/godotenv"
)

// Run Setup and runs an entire application
func Run() {
	/*
		SETUP
	*/

	// Import env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Read .env file error: %v", err)
	}

	// Create channel for interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create channels for websocket-bot comunication
	requestGenerateChannel := make(chan models.GeneratedImage)
	responseGenerateChannel := make(chan models.GeneratedImage)

	// Start the DB connection
	repositories.DB = postgres.Init()

	// Start the Bot
	botDone := make(chan struct{})
	go client.Init(&botDone, requestGenerateChannel, responseGenerateChannel)

	// Start the Websocket connection
	go websocket.Init(requestGenerateChannel, responseGenerateChannel)

	// Handle channels
	<-interrupt
	log.Println("Interrupt signal received. Initiating shutdown.")
	close(botDone)
	log.Println("Application shut down gracefully")
}
