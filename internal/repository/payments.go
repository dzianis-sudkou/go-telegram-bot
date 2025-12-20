package repositories

import "github.com/dzianis-sudkou/go-telegram-bot/internal/models"

// CREATE

// CreatePayment Creates new payment
func CreatePayment(payment *models.Payment) (err error) {
	err = DB.Create(payment).Error
	return
}
