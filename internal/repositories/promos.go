package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

func CreatePromo(promo *models.Promo) (err error) {
	err = DB.Create(promo).Error
	return
}

// READ
func GetPromo(code string) (promo models.Promo, err error) {
	err = DB.Table("promos").Where("code", code).First(&promo).Error
	return
}

// UPDATE
func UpdatePromo(promo *models.Promo) (err error) {
	err = DB.Save(promo).Error
	return
}
