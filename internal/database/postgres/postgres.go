package postgres

import (
	"encoding/json"
	"log"
	"os"

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
		&models.EnLocale{},
		&models.RuLocale{},
	)
	if err != nil {
		log.Fatalf("Table creation error: %v", err)
	}
}

func GenerateLocales(db *gorm.DB) {

	// Generate table with Locales [ru, en]
	jsonData, err := os.ReadFile("internal/repositories/locales/en.json")
	if err != nil {
		panic(err)
	}

	var localesMap map[string]string
	err = json.Unmarshal(jsonData, &localesMap)
	if err != nil {
		log.Printf("JSON unmarshal error: %v", err)
	}

	for key, value := range localesMap {
		locale := models.EnLocale{
			State: key,
			Text:  value,
		}
		db.Create(&locale)
	}
}
