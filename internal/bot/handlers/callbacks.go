package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/bot/keyboards"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	// Send the answer to the telegram server
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Request the Callback Error: %v", err)
	}

	var msg tgbotapi.EditMessageTextConfig

	callbackData := strings.Split(update.CallbackData(), `_`)

	switch callbackData[0] {

	case "start":
		msg = callbackStart(bot, &update)

	case "generate":
		msg = callbackGenerate(&update, &callbackData)

	case "request":
		msg = callbackRequest(&update, &callbackData)

	case "download":
		msg = callbackDownload(&update, &callbackData)

	case "socials":
		msg = callbackSocials(&update)

	case "payment":
		msg = callbackPayment(bot, &update, &callbackData)

	case "support":
		msg = callbackSupport(&update)

	}
	if msg.Text != "" {
		msg.ParseMode = "HTML"
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Callback Error: %v", err)
		}
	}
}

func callbackStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (msg tgbotapi.EditMessageTextConfig) {
	services.SetUserState(update, "start")
	if services.IsSubscribed(bot, update.CallbackQuery.From.ID) {
		text := fmt.Sprintf(services.GetTextLocale(update.CallbackQuery.From.LanguageCode, update.CallbackData()), update.CallbackQuery.From.FirstName+update.CallbackQuery.From.LastName)
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardStart(update.CallbackQuery.From.LanguageCode),
		)
	} else {
		text := "To use this bot, you should be subscribed to my channel!"
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardSubscribe(),
		)
	}
	return
}

func callbackGenerate(update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := (*callbackData)[0] + "_" + (*callbackData)[1]

	switch (*callbackData)[1] {

	// Menu that prints all the information and let's user choose the model
	case "menu":
		user := services.GetUser(update)
		services.SetUserState(update, state)
		text := fmt.Sprintf(
			services.GetTextLocale(update.SentFrom().LanguageCode, "generate_menu"),
			user.FullName,
			user.GeneratedImagesCount,
			user.Credits,
			uint(float64(user.Credits)*12.5),
		)

		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.FromChat().ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardGenerateMenu(),
		)

	case "1":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.FromChat().ID,
			update.CallbackQuery.Message.MessageID,
			"I'm sorry this feature is not available yet 😔\nCome back later",
			keyboards.KeyboardBackButton("generate_menu"),
		)

	case "2":
		if services.EnoughCoins(2, update) {
			services.SetUserState(update, state)
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				"Now send the description of the image in the realism style.\nAnd I'll create it for you.",
				keyboards.KeyboardBackButton("generate_menu"),
			)
		} else {
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				"You don't have enough coins. 🪙\nTop up your balance and come back here.",
				keyboards.KeyboardBackButton("generate_menu"),
			)
		}

	case "3":
		if services.EnoughCoins(2, update) {
			services.SetUserState(update, state)
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				"Now send the description of the image in the anime style.\nAnd I'll create it for you.",
				keyboards.KeyboardBackButton("generate_menu"),
			)
		} else {
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				"You don't have enough coins. 🪙\nTop up your balance and come back here.",
				keyboards.KeyboardBackButton("generate_menu"),
			)
		}
	}
	return
}

func callbackPayment(bot *tgbotapi.BotAPI, update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := (*callbackData)[0] + "_" + (*callbackData)[1]
	services.SetUserState(update, state)

	switch (*callbackData)[1] {

	// Generate image menu
	case "menu":
		text := services.GetTextLocale(update.SentFrom().LanguageCode, "payment_menu")
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.FromChat().ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardPayment(),
		)
	case "1000":
		createNewInvoice(bot, update, 1000)

	case "500":
		createNewInvoice(bot, update, 500)

	case "250":
		createNewInvoice(bot, update, 250)
	}
	return
}

func createNewInvoice(bot *tgbotapi.BotAPI, update *tgbotapi.Update, amount int) {

	prices := []tgbotapi.LabeledPrice{
		{
			Label:  fmt.Sprintf("Pay for %d stars", amount),
			Amount: amount,
		},
	}

	title := fmt.Sprintf(services.GetTextLocale(update.SentFrom().LanguageCode, "balance_up_title"), int(float64(amount)/12.5))
	description := fmt.Sprintf(services.GetTextLocale(update.SentFrom().LanguageCode, "balance_up_description"), int(float64(amount)/12.5))

	invoice := tgbotapi.NewInvoice(update.FromChat().ID, title, description, "custom_payload", "", "start_param", "XTR", prices)
	invoice.SuggestedTipAmounts = []int{}

	if _, err := bot.Send(invoice); err != nil {
		log.Printf("Sending the invoice error: %v", err)
	}
}

func callbackRequest(update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := (*callbackData)[0] + "_" + (*callbackData)[1]

	services.SetUserState(update, state)

	text := services.GetTextLocale(update.CallbackQuery.From.LanguageCode, state)

	switch (*callbackData)[1] {

	// Menu for choosing the type of request
	case "0":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardRequestTypes(),
		)

	// Paid Request type
	case "1":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardPaidPictureRequest(),
		)

	// Free Request type
	case "2":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardFreeRequestStart(),
		)

	// Request form
	case "make":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode),
		)

	}
	return
}

func callbackDownload(update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := (*callbackData)[0] + "_" + (*callbackData)[1]

	services.SetUserState(update, state)

	text := services.GetTextLocale(update.CallbackQuery.From.LanguageCode, state)

	switch (*callbackData)[1] {

	// Menu for the download
	case "0":
		msg = tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ID, update.CallbackQuery.Message.MessageID, text, keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode))
	}
	return
}

func callbackSocials(update *tgbotapi.Update) (msg tgbotapi.EditMessageTextConfig) {
	msg = tgbotapi.NewEditMessageTextAndMarkup(
		update.FromChat().ID,
		update.CallbackQuery.Message.MessageID,
		services.GetTextLocale(update.SentFrom().LanguageCode, "socials"),
		keyboards.KeyboardSocials(),
	)
	return
}

func callbackSupport(update *tgbotapi.Update) (msg tgbotapi.EditMessageTextConfig) {
	msg = tgbotapi.NewEditMessageTextAndMarkup(
		update.FromChat().ID,
		update.CallbackQuery.Message.MessageID,
		services.GetTextLocale(update.SentFrom().LanguageCode, "support"),
		keyboards.KeyboardSupport(),
	)
	return
}
