package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE
func CreateImage(image *models.Image) error {
	result := DB.Create(image)
	return result.Error
}

// READ
func GetImagesByPostID(postId uint) (img []models.Image, err error) {
	err = DB.Table("images").Where("post_id", postId).Scan(&img).Error
	return
}
