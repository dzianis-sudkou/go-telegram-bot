package keyboards

import (
	"fmt"
	"strings"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func KeyboardStart(locale string) tgbotapi.InlineKeyboardMarkup {
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "generateButton"), "generate_menu"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "requestButton"), "request_0"),
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "downloadButton"), "download_0"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "socialsButton"), "socials"),
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "supportButton"), "support"),
		),
	)
	return keyboard
}

func KeyboardSocials() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🟣 Instagram 🟣", "https://www.instagram.com/gokuryo_"),
			tgbotapi.NewInlineKeyboardButtonURL("⚫️ TikTok ⚫️", "https://www.tiktok.com/@gokuryo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🔴 Pinterest 🔴", "https://www.pinterest.com/gokuryo_"),
			tgbotapi.NewInlineKeyboardButtonURL("⚪️ Twitter ⚪️", "https://twitter.com/gokuryo_"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "start"),
		),
	)
	return keyboard
}

func KeyboardSupport() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🔵 PayPal 🔵", "https://www.paypal.com/donate/?hosted_button_id=R5C8W4VRS9Y8C"),
			tgbotapi.NewInlineKeyboardButtonURL("🟠 Boosty 🟠", "https://boosty.to/moskvinssss/donate"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Other...", "https://t.me/gokuryo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "start"),
		),
	)
	return keyboard
}

func KeyboardRequestTypes() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1️⃣", "request_1"),
			tgbotapi.NewInlineKeyboardButtonData("2️⃣", "request_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "start"),
		),
	)
	return keyboard
}

func KeyboardMainMenu(locale string) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(services.GetTextLocale(locale, "mainMenuButton"), "start"),
		),
	)
	return keyboard
}

func KeyboardSubscribe() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Creative Dream AI", "https://t.me/gokuryo_art"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅", "start"),
		),
	)
	return keyboard
}

func KeyboardPaidPictureRequest() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("PayPal", "https://www.paypal.com/donate/?hosted_button_id=R5C8W4VRS9Y8C"),
			tgbotapi.NewInlineKeyboardButtonURL("➡️", "https://t.me/gokuryo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "request_0"),
		),
	)
	return keyboard
}

func KeyboardFreeRequestStart() (keyboard tgbotapi.InlineKeyboardMarkup) {
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➡️", "request_make"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "request_0"),
		),
	)
	return
}

func KeyboardGenerateMenu(locale string) (keyboard tgbotapi.InlineKeyboardMarkup) {
	buttons := strings.Split(services.GetTextLocale(locale, "generate_menu_buttons"), "\n")

	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[0], "generate_anime_square_HD"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[1], "generate_realism_square_HD"),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData(buttons[2], "generate_creativedream_square_HD"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[3], "payment_menu"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[4], "start"),
		),
	)
	return
}

// generate_anime_square_HD
func KeyboardChooseFormat(model string, format string, quality string) (keyboard tgbotapi.InlineKeyboardMarkup) {

	formatInt := map[string]int{
		"horizontal": 0,
		"square":     1,
		"vertical":   2,
	}

	qualityInt := map[string]int{
		"HD": 0,
		"4K": 1,
	}

	data := "generate_%s_%s_%s"

	formatCallbackData := []string{
		fmt.Sprintf(data, model, "horizontal", quality),
		fmt.Sprintf(data, model, "square", quality),
		fmt.Sprintf(data, model, "vertical", quality),
	}

	qualityCallbackData := []string{
		fmt.Sprintf(data, model, format, "HD"),
		fmt.Sprintf(data, model, format, "4K"),
	}

	formatButtons := []tgbotapi.InlineKeyboardButton{
		{
			Text:         "16 : 9",
			CallbackData: &formatCallbackData[0],
		},
		{
			Text:         "1 : 1",
			CallbackData: &formatCallbackData[1],
		},
		{
			Text:         "9 : 16",
			CallbackData: &formatCallbackData[2],
		},
	}

	qualityButtons := []tgbotapi.InlineKeyboardButton{
		{
			Text:         "HD",
			CallbackData: &qualityCallbackData[0],
		},
		{
			Text:         "4K",
			CallbackData: &qualityCallbackData[1],
		},
	}

	formatButtons[formatInt[format]].Text = "✅ " + formatButtons[formatInt[format]].Text
	qualityButtons[qualityInt[quality]].Text = "✅ " + qualityButtons[qualityInt[quality]].Text
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			formatButtons...,
		),
		tgbotapi.NewInlineKeyboardRow(
			qualityButtons...,
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "generate_menu"),
		),
	)
	return
}

func KeyboardPayment() (keyboard tgbotapi.InlineKeyboardMarkup) {
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⭐️1000 (🪙 80)", "payment_1000"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⭐️500 (🪙 40)", "payment_500"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⭐️250 (🪙 20)", "payment_250"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "generate_menu"),
		),
	)
	return
}

func KeyboardBackButton(baskState string) (keyboard tgbotapi.InlineKeyboardMarkup) {
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", baskState),
		),
	)
	return
}

func KeyboardAcceptRules() (keyboard tgbotapi.InlineKeyboardMarkup) {
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅", "generate_acceptrules"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "start"),
		),
	)

	return
}
