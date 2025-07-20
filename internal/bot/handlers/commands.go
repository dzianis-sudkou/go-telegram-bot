package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {

	// /start
	case "start":
		msg = cmdStart(bot, update)

	// /downloadAllImages
	case "downloadAllImages":
		go cmdDownloadAllImages(bot, &update)

	// /addPost 100
	case "addPost":
		msg = cmdAddPost(&update)

	// /addCredits tgId amount
	case "addCredits":
		msg = cmdAddCredits(&update)

	// /addPromo code amount
	case "addPromo":
		msg = cmdAddPromo(&update)

	// /promo code
	case "promo":
		msg = cmdPromo(&update)

	default:
		msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "wrong_message"))
		msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)
	}

	removeLastMessage(bot, update.FromChat().ID, services.GetBotLastMessage(update.SentFrom().ID))

	msg.ParseMode = "HTML"
	lastMsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Sending the message error: %v", err)
	}
	services.UpdateLastMessage(update.SentFrom().ID, &lastMsg)
}

func cmdStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Add new user to the database if not presented
	services.AddNewUser(&update)

	// Control if the user is subscribed to the channel
	if services.IsSubscribed(bot, update.Message.From.ID) {
		services.SetUserState(&update, "start")
		// Getting the message reply from the locale database
		text := fmt.Sprintf(services.GetTextLocale(update.Message.From.LanguageCode, "start"), update.SentFrom().FirstName+update.SentFrom().LastName)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = keyboards.KeyboardStart(update.Message.From.LanguageCode)
	} else {
		text := fmt.Sprint(services.GetTextLocale(update.Message.From.LanguageCode, "not_subscribed"))
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = keyboards.KeyboardSubscribe()
	}
	return
}

func handlePrecheckoutQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	pca := tgbotapi.PreCheckoutConfig{
		OK:                 true,
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
	}

	if _, err := bot.Request(pca); err != nil {
		log.Printf("Answer preckout query: %v", err)
	}
}

func cmdDownloadAllImages(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	var mediaGroupDocs []any
	images := services.GetAllImages(update)
	for i, image := range images {
		if i%10 == 0 {
			time.Sleep(time.Second) // Avoid the bot api limits
			if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
				log.Printf("Send media group: %v", err)
			}
			mediaGroupDocs = make([]any, 0)
		}
		mediaGroupDocs = append(mediaGroupDocs, tgbotapi.NewInputMediaDocument(tgbotapi.FileID(image.ImageHash)))
	}
	if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
		log.Printf("Send media group: %v", err)
	}
	msg := tgbotapi.NewMessage(update.FromChat().ID, "All images were sent.")
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Sending the message error: %v", err)
	}
}

func cmdAddCredits(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Check if the user is admin
	if services.IsAdmin(update) {
		data := strings.Split(update.Message.CommandArguments(), " ") // [tgId, credits]
		if len(data) != 2 {
			msg = tgbotapi.NewMessage(update.FromChat().ID, "Wrong command")
			return
		}
		tgId, _ := strconv.ParseInt(data[0], 10, 64)
		amount, _ := strconv.Atoi(data[1])
		services.ChangeUserBalance(tgId, amount)
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Successful command execution - "+update.Message.Command())
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, you don't have access to this command.")
	}
	return
}

func cmdAddPost(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Check if the user is admin
	if services.IsAdmin(update) {
		services.SetUserState(update, "newPost_"+update.Message.CommandArguments()) // Change his state
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Now send me the description of your post #"+update.Message.CommandArguments())
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, you don't have access to this command.")
	}
	return
}

func cmdAddPromo(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Check if the user is admin
	if services.IsAdmin(update) {
		data := strings.Split(update.Message.CommandArguments(), " ") // [code, amount]
		if len(data) != 3 {
			msg = tgbotapi.NewMessage(update.FromChat().ID, "Wrong command")
			return
		}
		code := data[0]
		amount, _ := strconv.Atoi(data[1])
		activations, _ := strconv.Atoi(data[2])
		services.AddNewPromo(code, amount, activations)
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Successful command execution - "+update.Message.Command())
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, you don't have access to this command.")
	}
	return
}

func cmdPromo(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	promo := update.Message.CommandArguments()
	if promo == "" || !services.UsePromo(update, promo) {
		msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "wrong_promo"))
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "used_promo"))
	}
	msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)
	return
}
