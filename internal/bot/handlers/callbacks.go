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
			log.Printf("Send callback: %v", err)
		}
	}
}

func callbackStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (msg tgbotapi.EditMessageTextConfig) {
	services.SetUserState(update, "start")
	if services.IsSubscribed(bot, update.CallbackQuery.From.ID) {
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			fmt.Sprintf(services.GetTextLocale(update.SentFrom().LanguageCode, "start"), update.SentFrom().FirstName+update.SentFrom().LastName),
			keyboards.KeyboardStart(update.CallbackQuery.From.LanguageCode),
		)
	} else {
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			services.GetTextLocale(update.SentFrom().LanguageCode, "not_subscribed"),
			keyboards.KeyboardSubscribe(),
		)
	}
	return
}

func callbackGenerate(update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := update.CallbackData()
	log.Printf("Inside callback Generate - %s", state)
	stateSlice := getStateSlice(&state)

	switch stateSlice[1] {

	// Menu that prints all the information and let's user choose the model
	case "menu", "acceptrules":
		if stateSlice[1] == "acceptrules" {
			services.AcceptRules(update)
		}
		// Prints all the rules if the user hasn't accepted rules
		user := services.GetUser(update)
		if !user.Authorized {
			services.SetUserState(update, "generate_rules")
			text := services.GetTextLocale(update.SentFrom().LanguageCode, "generate_rules")
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				text,
				keyboards.KeyboardAcceptRules(),
			)
		} else {
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
				keyboards.KeyboardGenerateMenu(update.SentFrom().LanguageCode),
			)
		}

	case "anime", "realism": // generate_anime_square_HD
		var keyboard tgbotapi.InlineKeyboardMarkup
		keyboard = keyboards.KeyboardChooseFormat(stateSlice[1], stateSlice[2], stateSlice[3])
		if services.IsEnoughCoins(2, update) {
			services.SetUserState(update, state)
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,
				services.GetTextLocale(update.SentFrom().LanguageCode, "generate_prompt"),
				keyboard,
			)
		} else {
			msg = tgbotapi.NewEditMessageTextAndMarkup(
				update.FromChat().ID,
				update.CallbackQuery.Message.MessageID,

				services.GetTextLocale(update.SentFrom().LanguageCode, "insufficient_funds"),
				keyboards.KeyboardBackButton("generate_menu"),
			)
		}

	case "creativedream":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.FromChat().ID,
			update.CallbackQuery.Message.MessageID,
			services.GetTextLocale(update.SentFrom().LanguageCode, "not_available"),
			keyboards.KeyboardBackButton("generate_menu"),
		)
	}
	return
}

func callbackPayment(bot *tgbotapi.BotAPI, update *tgbotapi.Update, callbackData *[]string) (msg tgbotapi.EditMessageTextConfig) {
	state := (*callbackData)[0] + "_" + (*callbackData)[1]
	services.SetUserState(update, state)

	switch (*callbackData)[1] {

	// Payment menu
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

	invoice := tgbotapi.NewInvoice(update.FromChat().ID, title, description, "Coins for the image generation", "", "start_param", "XTR", prices)
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
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			update.FromChat().ID,
			update.CallbackQuery.Message.MessageID,
			text,
			keyboards.KeyboardMainMenu(update.SentFrom().LanguageCode),
		)
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
