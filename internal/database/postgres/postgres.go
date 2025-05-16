package postgres

import (
	"encoding/json"
	"fmt"
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

	// Generate the table with locales
	GenerateLocales(db)

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
	var languages []string = []string{"en", "ru"}

	for _, language := range languages {
		filePath := fmt.Sprintf("internal/repositories/locales/%s.json", language)

		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		var localesMap map[string]string
		err = json.Unmarshal(jsonData, &localesMap)
		if err != nil {
			log.Printf("JSON unmarshal error: %v", err)
		}
		for key, value := range localesMap {
			switch language {
			case "en":
				locale := models.EnLocale{
					State: key,
					Text:  value,
				}
				db.Create(&locale)
			case "ru":
				locale := models.RuLocale{
					State: key,
					Text:  value,
				}
				db.Create(&locale)
			}

		}
	}
}
