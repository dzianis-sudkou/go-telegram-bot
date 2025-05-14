package handlers

import (
	"fmt"
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	// Send the answer to the telegram server
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Request the Callback Error: %v", err)
	}

	var msg tgbotapi.EditMessageTextConfig

	switch update.CallbackData() {
	case "request":
		text := "<b>1️⃣ - (PAID)</b> Skip the queue and get your image as soon as possible.\n"
		text += "<b>2️⃣ - (FREE)</b> Your request will be added to the queue and will be processed.\n\n"
		text += "Choose a type of your request:"
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardRequestTypes())
	case "download":
	case "socials":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "<b>🔗 Socials 🔗</b>", keyboards.KeyboardSocials())
	case "support":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "<b>💵 Support Channel 💵</b>", keyboards.KeyboardSupport())
	case "start":
		text := fmt.Sprintf("<b>Hello %s</b>\n", update.SentFrom().FirstName+update.SentFrom().LastName)
		text += "This is the bot Creative Dream AI.\nHere you can:\n"
		text += "<b>1.</b> Make a request for your own character.\n"
		text += "<b>2.</b> Download my pictures without watermark in the best quality."
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardStart())
	}
	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Callback Error: %v", err)
	}
}
