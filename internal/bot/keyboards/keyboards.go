package keyboards

import (
	"github.com/dzianis-sudkou/go-telegram-bot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func KeyboardStart(locale string) tgbotapi.InlineKeyboardMarkup {
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
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

func KeyboardMainMenu() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🏠 Main Menu 🏠", "start"),
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

func KeyboardFreeRequestStart() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➡️", "request_make"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️", "request_0"),
		),
	)
	return keyboard
}
