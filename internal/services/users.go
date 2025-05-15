package services

import (
	"log"
	"slices"
	"strconv"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
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

func IsSubscribed(bot *tgbotapi.BotAPI, tgID int64) bool {

	// Take the channel id from env
	channelId, _ := strconv.ParseInt(config.Config("CHANNEL_ID"), 10, 64)
	memberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: channelId,
			UserID: tgID,
		},
	}

	member, err := bot.GetChatMember(memberConfig)
	if err != nil {
		log.Printf("Member is not presented: %v", err)
		return false
	}
	roles := []string{"creator", "administrator", "member"}
	return slices.Contains(roles, member.Status)
}
