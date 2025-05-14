package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func KeyboardStart() tgbotapi.InlineKeyboardMarkup {
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍️ Request ✍️", "request"),
			tgbotapi.NewInlineKeyboardButtonData("🖼️ Download 🖼️", "download"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔗 Socials 🔗", "socials"),
			tgbotapi.NewInlineKeyboardButtonData("💵 Support channel 💵", "support"),
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
			tgbotapi.NewInlineKeyboardButtonData("1️⃣", "paid_request_0"),
			tgbotapi.NewInlineKeyboardButtonData("2️⃣", "free_request_0"),
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
