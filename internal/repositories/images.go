package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE
func CreateImage(image *models.Image) (err error) {
	err = DB.Create(image).Error
	return
}

// READ
func GetImagesByPostID(postId uint) (img []models.Image, err error) {
	err = DB.Table("images").Where("post_id", postId).Scan(&img).Error
	return
}
