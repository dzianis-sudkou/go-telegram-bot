package services

import (
	"log"
	"strconv"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddNewPost(update *tgbotapi.Update, number string) {
	postId, err := strconv.Atoi(number)
	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}
	newPost := models.Post{
		ID:          uint(postId),
		Description: update.Message.Text,
	}
	if err := repositories.CreatePost(&newPost); err != nil {
		log.Printf("Post creation error: %v", err)
	}
}
