package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI) {
	// Drop updates
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	// Start polling telegram
	updates := bot.GetUpdatesChan(updateConfig)
	botInfo, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	log.Printf("Bot instance: @%s is online!", botInfo.UserName)

	// Go through each update update received from Telegram servers
	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message.IsCommand() {
			Commands(bot, update)
		} else if update.Message != nil {
			Messages(bot, update)
		} else {
			continue
		}
	}
}
