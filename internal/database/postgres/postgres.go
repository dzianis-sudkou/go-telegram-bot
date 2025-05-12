package postgres

import (
	"log"
	"os"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to database
func Connect() *gorm.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatalf("Environment Variable %s not found", "DSN")
	}
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
