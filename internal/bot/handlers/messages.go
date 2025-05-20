package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update, requestCh chan models.GeneratedImage) {
	var msg tgbotapi.MessageConfig

	state := services.GetUserState(&update)
	log.Printf("Message from User ID: %d", update.Message.MessageID)
	if update.Message.SuccessfulPayment != nil {
		log.Println(update.Message.SuccessfulPayment.TelegramPaymentChargeID, update.Message.SuccessfulPayment.TotalAmount)
		msg = msgSuccessfulPayment(&update)
	} else {
		stateSlice := strings.Split(state, "_")

		switch stateSlice[0] {

		// Admin has sent images for a new post
		case "newPost":
			msg = msgNewPost(&update, &stateSlice)

		// User asked to download images from the post
		case "download":
			msg = msgNewDownload(bot, &update)
			msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)

		// User has made a request
		case "request":
			if stateSlice[1] != "make" {
				msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "request_wrong"))
			} else {
				msg = msgNewRequest(&update)
			}
			msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)

		case "generate":
			msg = msgGenerate(&update, &stateSlice, requestCh)
		}
	}

	if msg.Text == "" {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, your request is wrong")
		msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)
	}
	msg.ParseMode = "HTML"
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
	if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
		log.Printf("Send media group: %v", err)
	}
	msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "download_1"))
	return
}

func msgNewRequest(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	services.AddNewRequest(update)
	text := fmt.Sprintf(services.GetTextLocale(update.SentFrom().LanguageCode, "request_made"), update.Message.Text)
	msg = tgbotapi.NewMessage(update.FromChat().ID, text)
	return
}

func msgSuccessfulPayment(update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	services.SetUserState(update, "start")
	services.AddNewPayment(update.Message.SuccessfulPayment)
	services.ChangeBalance(update.Message.SuccessfulPayment.TotalAmount, update)
	msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "successful_payment"))
	msg.ReplyMarkup = keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode)
	return
}

func msgGenerate(update *tgbotapi.Update, stateSlice *[]string, requestCh chan models.GeneratedImage) (msg tgbotapi.MessageConfig) {
	switch (*stateSlice)[1] {
	case "menu":
	case "anime", "realism", "creativedream":
		services.SetUserState(update, "start")
		switch (*stateSlice)[1] {
		case "anime":
			services.AddNewGeneratedImage(update, "anime", requestCh)
			services.ChangeBalance(-4, update)
		case "realism":
			services.AddNewGeneratedImage(update, "realism", requestCh)
			services.ChangeBalance(-2, update)
		case "creativedream":
			services.AddNewGeneratedImage(update, "creativedream", requestCh)
			services.ChangeBalance(-2, update)
		}

		msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "processing_generation"))
	}
	return
}
