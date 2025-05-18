package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI, botDone *chan struct{}) {
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
					Messages(bot, update)
				}
			} else if update.CallbackQuery != nil {
				Callbacks(bot, update)
			} else if update.PreCheckoutQuery != nil {
				handlePrecheckoutQuery(bot, &update)
			}
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
