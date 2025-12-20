package services

import (
	"log"
	"strconv"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewImage Adds new image to the table images
func AddNewImage(update *tgbotapi.Update, post string) {
	postID, err := strconv.ParseInt(post, 10, 0)
	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}

	newImage := models.Image{
		Name:      update.Message.Document.FileName,
		ImageHash: update.Message.Document.FileID,
		PostID:    uint(postID),
	}

	if err := repositories.CreateImage(&newImage); err != nil {
		log.Printf("Image creation error: %v", err)
	}
}

// GetImagesByPostID Gets all images with the passed postID
func GetImagesByPostID(update *tgbotapi.Update, post string) (images []models.Image) {
	postID, err := strconv.ParseInt(post, 10, 0)
	if err != nil {
		log.Printf("Post id is wrong: %v", err)
	}

	images, err = repositories.GetImagesByPostID(uint(postID))
	if err != nil {
		log.Printf("Getting images from database: %v", err)
	}
	return images
}

// GetAllImages Gets all images from the table images
func GetAllImages(update *tgbotapi.Update) (images []models.Image) {
	images, err := repositories.GetAllImages()
	if err != nil {
		log.Printf("Get all images: %v", err)
	}
	return
}
