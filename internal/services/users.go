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
		ChatId:               update.FromChat().ID,
		TgId:                 update.SentFrom().ID,
		FullName:             update.SentFrom().FirstName + update.SentFrom().LastName,
		MsgCount:             0,
		FreeRequestCount:     0,
		Credits:              0,
		GeneratedImagesCount: 0,
		RegistrationDate:     time.Now(),
		State:                "start",
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

func GetUserState(update *tgbotapi.Update) string {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
		return ""
	}
	return user.State
}

func SetUserState(update *tgbotapi.Update, state string) {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.State = state
	err = repositories.UpdateUser(&user)
	if err != nil {
		log.Printf("Update user state: %v", err)
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

func IsAdmin(update *tgbotapi.Update) bool {
	adminList := []string{
		config.Config("TG_GOKURYO_ID"),
		config.Config("TG_DZIANIS_ID"),
	}
	return slices.Contains(adminList, strconv.Itoa(int(update.SentFrom().ID)))
}

func GetUser(update *tgbotapi.Update) (user models.User) {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	return
}

func ChangeBalance(amount int, update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.Credits += amount
	if err := repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

func IsEnoughCoins(amount int, update *tgbotapi.Update) bool {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	return user.Credits >= amount
}

func UpdateMessageCount(update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.MsgCount += 1
	if err = repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

func AcceptRules(update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.Authorized = true
	if err = repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}
