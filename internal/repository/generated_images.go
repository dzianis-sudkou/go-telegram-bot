package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

// CreateGeneratedImage Adds new Generated Images
func CreateGeneratedImage(image *models.GeneratedImage) (err error) {
	err = DB.Create(image).Error
	return
}

// READ

// GetGeneratedImageByUUID Retrieves the generated image by UUID
func GetGeneratedImageByUUID(uuid string) (image models.GeneratedImage, err error) {
	err = DB.Table("generated_images").Where("task_uuid", uuid).First(&image).Error
	return
}

// UPDATE

// UpdateGeneratedImage Update the generated image info
func UpdateGeneratedImage(image models.GeneratedImage) (err error) {
	err = DB.Save(&image).Error
	return
}
