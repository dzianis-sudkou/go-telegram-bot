package services

import (
	"fmt"

	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
)

// GetTextLocale Gets the string in the user's locale
func GetTextLocale(locale string, state string) string {
	var (
		text string
		err  error
	)
	switch locale {
	// User's language is Russian
	case "ru":
		text, err = repositories.GetTextLocale("ru_locales", state)
	// User's language is English or anything else
	default:
		text, err = repositories.GetTextLocale("en_locales", state)
	}
	if err != nil {
		fmt.Printf("Get Text Locale Error: %v", err)
	}
	return text
}
