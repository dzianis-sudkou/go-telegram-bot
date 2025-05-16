package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var msg tgbotapi.MessageConfig

	state := services.GetUserState(&update)
	stateSlice := strings.Split(state, "_")

	switch stateSlice[0] {
	case "newPost":
		msg = msgNewPost(&update, &stateSlice)
	case "download":
		msg = msgNewDownload(bot, &update)
	default:

	}

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Message sending error: %v", err)
	}
}

func msgNewPost(update *tgbotapi.Update, stateSlice *[]string) (msg tgbotapi.MessageConfig) {
	if update.Message.Text != "" {
		services.AddNewPost(update, (*stateSlice)[1])
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Successfully created post #"+(*stateSlice)[1]+"\nNow send images")
	} else {
		services.AddNewImage(update, (*stateSlice)[1])
		msg = tgbotapi.NewMessage(update.FromChat().ID, "✅ - New image: "+update.Message.Document.FileName)
	}
	return
}

func msgNewDownload(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	var mediaGroupDocs []any
	images := services.GetImagesByPostId(update, update.Message.Text)
	for _, image := range images {
		mediaGroupDocs = append(mediaGroupDocs, tgbotapi.NewInputMediaDocument(tgbotapi.FileID(image.ImageID)))
	}
	fmt.Printf("%v", mediaGroupDocs)
	if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
		log.Printf("Send media group: %v", err)
	}
	msg = tgbotapi.NewMessage(update.FromChat().ID, "To proceed, press the button below.")
	return
}
