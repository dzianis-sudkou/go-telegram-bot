package services

import (
	"log"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
