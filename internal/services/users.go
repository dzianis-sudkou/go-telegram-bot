package services

import (
	"log"
	"slices"
	"strconv"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewUser Adds new user to the db
func AddNewUser(update *tgbotapi.Update) {
	newUser := models.User{
		ChatId:               update.FromChat().ID,
		TgId:                 update.SentFrom().ID,
		FullName:             update.SentFrom().FirstName + update.SentFrom().LastName,
		MsgCount:             0,
		FreeRequestCount:     0,
		Credits:              20,
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

// GetUserState Gets the user's state
func GetUserState(update *tgbotapi.Update) string {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
		return ""
	}
	return user.State
}

// SetUserState Changes the user's state in the menu.
func SetUserState(update *tgbotapi.Update, state string) {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.State = state
	err = repositories.UpdateUser(&user)
	if err != nil {
		log.Printf("Update user state: %v", err)
	}
}

// IsSubscribed Checks the user's subscription to the channel.
func IsSubscribed(bot *tgbotapi.BotAPI, tgID int64) bool {
	// Take the channel id from env
	channelID, _ := strconv.ParseInt(config.Config("CHANNEL_ID"), 10, 64)
	memberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: channelID,
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

// IsAdmin Checks if the user is Admin
func IsAdmin(update *tgbotapi.Update) bool {
	adminList := []string{
		config.Config("TG_GOKURYO_ID"),
		config.Config("TG_DZIANIS_ID"),
	}
	return slices.Contains(adminList, strconv.Itoa(int(update.SentFrom().ID)))
}

// GetUser Retrieves the user from database
func GetUser(update *tgbotapi.Update) (user models.User) {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	return
}

// ChangeBalance Changes the user balance by the desired amount
func ChangeBalance(amount int, update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.Credits += amount
	if err := repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

// ChangeUserBalance Changes the user balance
func ChangeUserBalance(tgID int64, amount int) {
	user, err := repositories.GetUserByTgID(tgID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.Credits += amount
	if err := repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

// IsEnoughCoins Checks if user has enough coins
func IsEnoughCoins(amount int, update *tgbotapi.Update) bool {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	return user.Credits >= amount
}

// UpdateMessageCount Updates user message count
func UpdateMessageCount(update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.MsgCount++
	if err = repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

// AcceptRules Updates the user information after the rules accept
func AcceptRules(update *tgbotapi.Update) {
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		log.Printf("User not found: %v", err)
	}
	user.Authorized = true
	if err = repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}
}

// UpdateLastMessage Updates the count of user messages + the id of bot message
func UpdateLastMessage(tgID int64, lastMsg *tgbotapi.Message) {
	user, err := repositories.GetUserByTgID(tgID)
	if err != nil {
		log.Printf("Get user: %v", err)
	}
	user.MsgCount++
	user.BotMessageID = lastMsg.MessageID
	if err := repositories.UpdateUser(&user); err != nil {
		log.Printf("BotMessageID update: %v", err)
	}
}

// GetBotLastMessage Retrieves the id of the last bot message using the user ID
func GetBotLastMessage(tgID int64) int {
	user, err := repositories.GetUserByTgID(tgID)
	if err != nil {
		log.Printf("Get user: %v", err)
	}
	return user.BotMessageID
}
