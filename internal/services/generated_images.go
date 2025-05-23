package services

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func AddNewGeneratedImage(update *tgbotapi.Update, model string, format string, quality string, requestCh chan models.GeneratedImage) {

	// Update user's number of generated images
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		log.Printf("Get user by TG id: %v", err)
	}
	user.GeneratedImagesCount += 1
	if err = repositories.UpdateUser(&user); err != nil {
		log.Printf("Update user: %v", err)
	}

	image := models.GeneratedImage{
		TaskType: "imageInference",
		Message:  int64(update.Message.MessageID) + 1,
		Prompt:   update.Message.Text,
		TaskUUID: uuid.NewString(),
		Done:     false,
		Chat:     update.FromChat().ID,
		Model:    model,
		Format:   format,
		Quality:  quality,
		Language: update.SentFrom().LanguageCode,
	}

	if err := repositories.CreateGeneratedImage(&image); err != nil {
		log.Printf("Save generated image: %v", err)
	}
	requestCh <- image
}

func UpdateGeneratedImage(image *models.GeneratedImage) (img models.GeneratedImage) {
	img, err := repositories.GetGeneratedImageByUUID(image.TaskUUID)
	if err != nil {
		log.Printf("Get generated image by UUID: %v", err)
	}
	{
		img.TaskType = image.TaskType
		img.Done = true
		img.NSFW = image.NSFW
		img.ImageURL = image.ImageURL
	}
	if err = repositories.UpdateGeneratedImage(img); err != nil {
		log.Printf("Update generated image in db: %v", err)
	}
	return
}
