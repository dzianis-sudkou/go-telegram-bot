package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// Create
func CreatePost(post *models.Post) error {
	result := DB.Create(post)
	return result.Error
}
