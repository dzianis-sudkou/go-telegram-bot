package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

func CreateUsedPromo(usedPromo *models.UsedPromo) (err error) {
	err = DB.Create(usedPromo).Error
	return
}
