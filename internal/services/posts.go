package services

import (
	"log"
	"strconv"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewPost Add New Post Function
func AddNewPost(update *tgbotapi.Update, number string) {
	postID, err := strconv.ParseInt(number, 10, 0)
	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}
	newPost := models.Post{
		ID:          uint(postID),
		Description: update.Message.Text,
	}
	if err := repositories.CreatePost(&newPost); err != nil {
		log.Printf("Post creation error: %v", err)
	}
}
