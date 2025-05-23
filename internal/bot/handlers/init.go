package handlers

import (
	"log"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI, botDone *chan struct{}, requestCh chan models.GeneratedImage, responseCh chan models.GeneratedImage) {

	// Drop updates
	updateConfig := newUpdateConfig(bot)

	// Start polling telegram
	updates := bot.GetUpdatesChan(updateConfig)
	botInfo, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	log.Printf("Bot instance: @%s is online!", botInfo.UserName)

	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				if update.Message.IsCommand() {
					Commands(bot, update)
				} else {
					Messages(bot, update, requestCh)
				}
			} else if update.CallbackQuery != nil {
				Callbacks(bot, update)
			} else if update.PreCheckoutQuery != nil {
				handlePrecheckoutQuery(bot, &update)
			}
		case image := <-responseCh:
			sendGeneratedImage(bot, image)
		case <-*botDone:
			bot.StopReceivingUpdates()
			return
		}
	}
}

func newUpdateConfig(bot *tgbotapi.BotAPI) (updateConfig tgbotapi.UpdateConfig) {
	update, _ := bot.GetUpdates(tgbotapi.NewUpdate(0))
	if len(update) != 0 {
		log.Println("Pending Updates: ", len(update))
		updateConfig = tgbotapi.NewUpdate(update[len(update)-1].UpdateID + 1)
	} else {
		updateConfig = tgbotapi.NewUpdate(0)
	}
	updateConfig.Timeout = 30
	updateConfig.AllowedUpdates = []string{"message", "callback_query", "pre_checkout_query", "shipping_query", "chat_member"}
	return
}

func sendGeneratedImage(bot *tgbotapi.BotAPI, image models.GeneratedImage) {

	removeLastMessage(bot, image.Chat, int(image.Message))

	// Check if image violates bot rules
	if image.NSFW {
		msg := tgbotapi.NewMessage(image.Chat, services.GetTextLocale(image.Language, "detected_nsfw"))
		msg.ParseMode = "HTML"
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Send nsfw-error message: %v", err)
		}
	} else {
		file := tgbotapi.FileURL(image.ImageURL)
		switch image.Quality {

		// HD image
		case "HD":
			msg := tgbotapi.NewPhoto(image.Chat, file)
			msg.Caption = services.GetTextLocale(image.Language, "completed_generation")
			msg.ParseMode = "HTML"
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Send the file: %v", err)
			}

		// 4K image
		case "4K":
			msg := tgbotapi.NewDocument(image.Chat, file)
			msg.Caption = services.GetTextLocale(image.Language, "completed_generation")
			msg.ParseMode = "HTML"
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Send the file: %v", err)
			}
		}
	}

	newMsg := tgbotapi.NewMessage(image.Chat, services.GetTextLocale(image.Language, "after_completed_generation"))
	newMsg.ReplyMarkup = keyboards.KeyboardBackButton("generate_menu")
	newMsg.ParseMode = "HTML"

	lastMsg, err := bot.Send(newMsg)
	if err != nil {
		log.Printf("Send the Generate Menu Message: %v", err)
	}
	services.UpdateLastMessage(lastMsg.Chat.ID, &lastMsg)
}

func getStateSlice(state *string) (stateSlice []string) {
	stateSlice = strings.Split(*state, "_")
	log.Printf("Get state Slice: %+v     %s", stateSlice, *state)
	return
}

// Remove the previous message
func removeLastMessage(bot *tgbotapi.BotAPI, chatId int64, messageId int) {
	deleteMessage := tgbotapi.NewDeleteMessage(chatId, messageId)
	if _, err := bot.Request(deleteMessage); err != nil {
		log.Printf("Delete Message: %v", err)
	}
}
