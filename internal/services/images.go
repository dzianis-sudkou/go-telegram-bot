package services

import (
	"log"
	"strconv"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddNewImage(update *tgbotapi.Update, post string) {
	postId, err := strconv.Atoi(post)

	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}

	newImage := models.Image{
		Name:    update.Message.Document.FileName,
		ImageID: update.Message.Document.FileID,
		PostID:  uint(postId),
	}

	if err := repositories.CreateImage(&newImage); err != nil {
		log.Printf("Image creation error: %v", err)
	}
}

func GetImagesByPostId(update *tgbotapi.Update, post string) (images []models.Image) {
	postId, err := strconv.Atoi(post)
	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}

	images, err = repositories.GetImagesByPostID(uint(postId))
	if err != nil {
		log.Printf("Getting images from database: %v", err)
	}
	return images
}

func GetAllImages(update *tgbotapi.Update) (images []models.Image) {
	images, err := repositories.GetAllImages()
	if err != nil {
		log.Printf("Get all images: %v", err)
	}
	return
}
