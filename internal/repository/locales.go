package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// GetTextLocale Gets locale string
func GetTextLocale(table string, state string) (string, error) {
	var text models.EnLocale
	result := DB.Table(table).Where("state", state).First(&text)
	return text.Text, result.Error
}
