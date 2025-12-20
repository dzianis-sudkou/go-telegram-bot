package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// Create

// CreatePost Creates new post
func CreatePost(post *models.Post) error {
	result := DB.Create(post)
	return result.Error
}
