package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

// CreateImage Creates new image
func CreateImage(image *models.Image) (err error) {
	err = DB.Create(image).Error
	return
}

// READ

// GetImagesByPostID Gets all images in the post
func GetImagesByPostID(postID uint) (img []models.Image, err error) {
	err = DB.Table("images").Where("post_id", postID).Scan(&img).Error
	return
}

// GetAllImages Gets all images
func GetAllImages() (img []models.Image, err error) {
	err = DB.Table("images").Scan(&img).Error
	return
}
