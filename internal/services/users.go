package services

import (
	"log"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddNewUser(update *tgbotapi.Update) {
	var newUser = models.User{
		TgId:                 uint64(update.Message.From.ID),
		FullName:             update.Message.From.FirstName + update.Message.From.LastName,
		MsgCount:             0,
		FreeRequestCount:     0,
		GeneratedImagesCount: 0,
		RegistrationDate:     time.Now(),
		State:                "null",
		Authorized:           false,
	}

	// Check if the user is already presented in the database
	if !repositories.IsLogged(&newUser) {
		err := repositories.CreateUser(&newUser)
		if err != nil {
			log.Printf("User creation error: %v", err)
		}
	}
}
