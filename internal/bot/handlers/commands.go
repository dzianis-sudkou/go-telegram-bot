package handlers

import (
	"fmt"
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		cmdStart(bot, update)
	case "help":
		cmdHelp(bot, update)
	}
}

func cmdStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var msg tgbotapi.MessageConfig

	// Control if the user is subscribed to the channel
	if services.IsSubscribed(bot, update.Message.From.ID) {

		// Add new user to the database if not presented
		services.AddNewUser(&update)

		text := fmt.Sprintf("<b>Hello %s</b>\n", update.SentFrom().FirstName+update.SentFrom().LastName)
		text += "This is the bot <b>Creative Dream AI</b>.\nHere you can:\n"
		text += "<b>1.</b> Make a request for your own character.\n"
		text += "<b>2.</b> Download my pictures without watermark in the best quality."
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = keyboards.KeyboardStart()
	} else {
		text := "To use this bot, you should be subscribed to my channel!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = keyboards.KeyboardSubscribe()
	}

	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Sending the message error: %v", err)
	}
}

func cmdHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "I don't have any help info yet...😔"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Sending the message error: %v", err)
	}
}
