package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	log.Println("Message Sent!")
}
