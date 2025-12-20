package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CreateUsedPromo Creates used promo code (idk why I have it 🤷‍♂️)
func CreateUsedPromo(usedPromo *models.UsedPromo) (err error) {
	err = DB.Create(usedPromo).Error
	return
}
