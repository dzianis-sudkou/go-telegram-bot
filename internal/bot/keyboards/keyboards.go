package keyboards

import (
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
			tgbotapi.NewInlineKeyboardButtonURL("Other...", "t.me/@gokuryo"),
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
			tgbotapi.NewInlineKeyboardButtonURL("Creative Dream AI", "t.me/@gokuryo_art"),
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
			tgbotapi.NewInlineKeyboardButtonURL("➡️", "t.me/@gokuryo"),
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
			tgbotapi.NewInlineKeyboardButtonData(buttons[0], "generate_1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[1], "generate_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[2], "generate_3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[3], "payment_menu"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttons[4], "start"),
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
