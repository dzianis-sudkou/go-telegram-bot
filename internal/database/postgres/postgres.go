package postgres

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Connect to database
	db := Connect()

	// Create Tables
	CreateTables(db)
	return db
}

// Connect to database
func Connect() *gorm.DB {
	dsn := config.Config("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Open db connection error: %v", err)
	}
	return db
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Image{},
		&models.FreeRequest{},
		&models.GeneratedImage{},
	)
	if err != nil {
		log.Fatalf("Table creation error: %v", err)
	}
}
