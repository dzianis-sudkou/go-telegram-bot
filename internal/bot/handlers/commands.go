package handlers

import (
	"fmt"
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {

	case "start":
		msg = cmdStart(bot, update)

	case "testGenerate":
		log.Println("Balance were changed!")
		services.ChangeBalance(100, &update)

	case "downloadAllImages":
		cmdDownloadAllImages(bot, &update)

	case "newPost":
		msg = cmdNewPost(update)
	}

	if msg.Text != "" {
		msg.ParseMode = "HTML"
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Sending the message error: %v", err)
		}
	}
}

func cmdStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Control if the user is subscribed to the channel
	if services.IsSubscribed(bot, update.Message.From.ID) {

		// Add new user to the database if not presented
		services.AddNewUser(&update)
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

func cmdNewPost(update tgbotapi.Update) (msg tgbotapi.MessageConfig) {

	// Check if the user is admin
	if services.IsAdmin(&update) {
		services.SetUserState(&update, "newPost_"+update.Message.CommandArguments()) // Change his state
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Now send me the description of your post #"+update.Message.CommandArguments())
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, you don't have access to this command.")
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

func cmdDownloadAllImages(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	var mediaGroupDocs []any
	images := services.GetAllImages(update)
	for i, image := range images {
		if i%10 == 0 {
			if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
				log.Printf("Send media group: %v", err)
			}
			mediaGroupDocs = make([]any, 0)
		}
		mediaGroupDocs = append(mediaGroupDocs, tgbotapi.NewInputMediaDocument(tgbotapi.FileID(image.ImageID)))
	}
	if _, err := bot.SendMediaGroup(tgbotapi.MediaGroupConfig{ChatID: update.FromChat().ID, Media: mediaGroupDocs}); err != nil {
		log.Printf("Send media group: %v", err)
	}
	msg = tgbotapi.NewMessage(update.FromChat().ID, services.GetTextLocale(update.SentFrom().LanguageCode, "download_1"))
	return
}
