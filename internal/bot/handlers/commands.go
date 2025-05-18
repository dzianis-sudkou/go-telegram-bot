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
	case "help":
		cmdHelp(bot, update)
	case "newPost":
		msg = cmdNewPost(bot, update)
	case "payment":
		cmdNewInvoice(bot, update)
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
		services.SetUserState(&update, "null")

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

func cmdHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var mediaGroupDocs []any
	files := []string{
		"BQACAgIAAxkBAAIXRGgnfn0XPJKzrpreGZ6n1mALkWsiAAIxcwACRBs5SZyy04QDPFfxNgQ",
		"BQACAgIAAxkBAAIXSWgnf81sMh6qqz0rNBqTXOaOGJT_AAJEcwACRBs5SX-BHZ0gApinNgQ",
		"BQACAgIAAxkBAAIXSmgnf81N1FpuRK5RSDnX7fxXkXIrAAJGcwACRBs5SbYi3xJeaUfFNgQ",
		"BQACAgIAAxkBAAIXS2gnf81Apxd6CWFFLRlj7zXzRu8OAAJHcwACRBs5Sa-e7yi7v-cwNgQ",
		"BQACAgIAAxkBAAIXTGgnf82VC3f0Whn45cbXTcB7TbGGAAJKcwACRBs5SXyOTIeYziiFNgQ",
		"BQACAgIAAxkBAAIXRGgnfn0XPJKzrpreGZ6n1mALkWsiAAIxcwACRBs5SZyy04QDPFfxNgQ",
	}

	for _, val := range files {
		mediaGroupDocs = append(mediaGroupDocs, tgbotapi.NewInputMediaDocument(tgbotapi.FileID(val)))
	}

	mediaGroup := tgbotapi.MediaGroupConfig{
		ChatID: update.Message.Chat.ID,
		Media:  mediaGroupDocs,
	}

	_, err := bot.SendMediaGroup(mediaGroup)
	if err != nil {
		log.Panicf("Error sending the mediagroup: %v", err)
	} else {
		log.Println("MediaGroup was sent successfully.")
	}
}

func cmdNewPost(bot *tgbotapi.BotAPI, update tgbotapi.Update) (msg tgbotapi.MessageConfig) {
	// Check if the user is admin
	if services.IsAdmin(&update) {
		services.SetUserState(&update, "newPost_"+update.Message.CommandArguments()) // Change his state
		msg = tgbotapi.NewMessage(update.FromChat().ID, "Now send me the description of your post #"+update.Message.CommandArguments())
	} else {
		msg = tgbotapi.NewMessage(update.FromChat().ID, "I'm sorry, you don't have access to this command.")
	}
	return
}

func cmdNewInvoice(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	prices := []tgbotapi.LabeledPrice{
		{
			Label:  "Pay for 10 credits",
			Amount: 1,
		},
	}

	invoice := tgbotapi.NewInvoice(update.FromChat().ID, "Test invoice", "description here", "custom_payload", "", "start_param", "XTR", prices)
	invoice.SuggestedTipAmounts = []int{}

	if _, err := bot.Send(invoice); err != nil {
		log.Printf("Sending the invoice error: %v", err)
	}
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
