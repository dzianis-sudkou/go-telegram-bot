package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

// CreatePromo Create new promo code
func CreatePromo(promo *models.Promo) (err error) {
	err = DB.Create(promo).Error
	return
}

// READ

// GetPromo Get the promo object
func GetPromo(code string) (promo models.Promo, err error) {
	err = DB.Table("promos").Where("code", code).First(&promo).Error
	return
}

// UPDATE

// UpdatePromo Updates the promo params
func UpdatePromo(promo *models.Promo) (err error) {
	err = DB.Save(promo).Error
	return
}
