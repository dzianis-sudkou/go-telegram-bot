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
	case "request":
		text := "<b>1️⃣ - (PAID)</b> Skip the queue and get your image as soon as possible.\n"
		text += "<b>2️⃣ - (FREE)</b> Your request will be added to the queue and will be processed.\n\n"
		text += "Choose a type of your request:"
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardRequestTypes())
	case "download":
		if len(callbackData) == 1 {
			text := "<b>Download</b>\n\nTo Download my pictures in the best quality and without watermark, "
			text += "send me a number of publication.\n\nHow to knew a number you need:\n"
			text += "<b>1.</b> Open publication in my group - @gokuryo_art\n"
			text += "<b>2.</b> Copy the number from the description."
			msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardMainMenu())
		} else {
			log.Println("This is the next step")
		}
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
	case "startVerify":
		if services.IsSubscribed(bot, update.CallbackQuery.From.ID) {
			text := fmt.Sprintf("<b>Hello %s</b>\n", update.SentFrom().FirstName+update.SentFrom().LastName)
			text += "This is the bot Creative Dream AI.\nHere you can:\n"
			text += "<b>1.</b> Make a request for your own character.\n"
			text += "<b>2.</b> Download my pictures without watermark in the best quality."
			msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardStart())
		} else {
			text := "To use this bot, you should be subscribed to my channel!"
			msg = tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardSubscribe())
		}
	}

	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Callback Error: %v", err)
	}
}
