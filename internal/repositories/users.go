package repositories

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"gorm.io/gorm"
)

// CREATE

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

// READ

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := DB.Table("users").Find(&users)
	return users, result.Error
}

func GetUserByTgId(id int64) (user models.User, err error) {
	err = DB.Table("users").Where("tg_id", id).First(&user).Error
	return
}

// UPDATE

func UpdateUser(user *models.User) error {
	result := DB.Save(user)
	return result.Error
}
