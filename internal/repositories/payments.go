package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

func CreatePayment(payment *models.Payment) (err error) {
	err = DB.Create(payment).Error
	return
}
