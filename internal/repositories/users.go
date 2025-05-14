package repositories

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"gorm.io/gorm"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := DB.Table("users").Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := DB.Table("users").Where("id", 1).First(&user)
	return user, result.Error
}

func CreateUser(user *models.User) error {
	result := DB.Create(user)
	return result.Error
}

func IsLogged(user *models.User) bool {
	var foundUser models.User
	err := DB.Table("users").Where("tg_id", user.TgId).Limit(1).First(&foundUser).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err == nil {
		if foundUser.TgId == user.TgId {
			return true
		}
	}
	log.Printf("Searching for the user error: %v", err)
	return false
}
