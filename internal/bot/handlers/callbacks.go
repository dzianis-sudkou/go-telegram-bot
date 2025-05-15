package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	// Send the answer to the telegram server
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Request the Callback Error: %v", err)
	}

	var msg tgbotapi.EditMessageTextConfig

	callbackData := strings.Split(update.CallbackData(), `_`)

	switch callbackData[0] {
	case "start":
		if services.IsSubscribed(bot, update.CallbackQuery.From.ID) {
			text := fmt.Sprintf(services.GetTextLocale(update.CallbackQuery.From.LanguageCode, update.CallbackData()), update.CallbackQuery.From.FirstName+update.CallbackQuery.From.LastName)
			msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardStart(update.CallbackQuery.From.LanguageCode))
		} else {
			text := "To use this bot, you should be subscribed to my channel!"
			msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardSubscribe())
		}

	case "request":
		msg = callbackRequest(&update, &callbackData)

	case "download":
		msg = callbackDownload(&update, &callbackData)

	case "socials":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, services.GetTextLocale(update.CallbackQuery.From.LanguageCode, "socials"), keyboards.KeyboardSocials())

	case "support":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, services.GetTextLocale(update.CallbackQuery.From.LanguageCode, "support"), keyboards.KeyboardSupport())

	}

	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Callback Error: %v", err)
	}
}

func callbackRequest(update *tgbotapi.Update, callbackData *[]string) tgbotapi.EditMessageTextConfig {
	var msg tgbotapi.EditMessageTextConfig

	services.SetUserState(update, (*callbackData)[0]+"_"+(*callbackData)[1])

	text := services.GetTextLocale(update.CallbackQuery.From.LanguageCode, (*callbackData)[0]+"_"+(*callbackData)[1])

	switch (*callbackData)[1] {

	// Menu for choosing the type of request
	case "0":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardRequestTypes())

	// Paid Request type
	case "1":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardPaidPictureRequest())

	// Free Request type
	case "2":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardFreeRequestStart())

	// Request a character form
	case "make":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardMainMenu())

	}
	return msg
}

func callbackDownload(update *tgbotapi.Update, callbackData *[]string) tgbotapi.EditMessageTextConfig {
	var msg tgbotapi.EditMessageTextConfig

	services.SetUserState(update, (*callbackData)[0]+"_"+(*callbackData)[1])

	text := services.GetTextLocale(update.CallbackQuery.From.LanguageCode, (*callbackData)[0]+"_"+(*callbackData)[1])

	switch (*callbackData)[1] {

	// Menu for the download
	case "0":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardMainMenu())
	}
	return msg
}
