package client

import (
	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/handlers"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() {
	bot, err := tgbotapi.NewBotAPI(config.Config("TG_API"))
	if err != nil {
		panic(err)
	}
	bot.Debug = config.BotDebug
	handlers.Init(bot)
}
