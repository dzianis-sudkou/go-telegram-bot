package repositories

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"gorm.io/gorm"
)

// CREATE

// CreateUser Adds new User to the database
func CreateUser(user *models.User) (err error) {
	err = DB.Create(user).Error
	return
}

// IsLogged Verifies if the user is logged
func IsLogged(user *models.User) bool {
	var foundUser models.User
	err := DB.Table("users").Where("tg_id", user.TgId).Limit(1).First(&foundUser).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return false
	case nil:
		if foundUser.TgId == user.TgId {
			return true
		}
	}
	log.Printf("Searching for the user error: %v", err)
	return false
}

// READ

// GetAllUsers retrieves all users from the database.
func GetAllUsers() (users []models.User, err error) {
	err = DB.Table("users").Find(&users).Error
	return
}

// GetUserByTgID retrieves the user by his id from database
func GetUserByTgID(id int64) (user models.User, err error) {
	err = DB.Table("users").Where("tg_id", id).First(&user).Error
	return
}

// UPDATE

// UpdateUser updates the user information on database
func UpdateUser(user *models.User) (err error) {
	err = DB.Save(user).Error
	return
}
