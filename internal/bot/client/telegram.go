package client

import (
	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/handlers"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Init the Telegram bot
func Init(botDone *chan struct{}, requestCh chan models.GeneratedImage, responseCh chan models.GeneratedImage) {
	bot, err := tgbotapi.NewBotAPI(config.Config("TG_API"))
	if err != nil {
		panic(err)
	}
	bot.Debug = config.BotDebug

	handlers.Init(bot, botDone, requestCh, responseCh)
}
