package services

import (
	"log"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewRequest adds new request to the database
func AddNewRequest(update *tgbotapi.Update) {
	request := models.FreeRequest{
		Text:         update.Message.Text,
		CreationDate: time.Now(),
		Language:     update.SentFrom().LanguageCode,
		UserID:       GetUser(update).ID,
	}
	if err := repositories.CreateRequest(&request); err != nil {
		log.Printf("Create request: %v", err)
	}
}
