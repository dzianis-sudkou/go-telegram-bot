package services

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewPromo Function too add a new Promo
func AddNewPromo(code string, amount int, activations int) {
	newPromo := models.Promo{
		Code:        code,
		Amount:      amount,
		UseCount:    0,
		Activations: activations,
	}

	err := repositories.CreatePromo(&newPromo)
	if err != nil {
		log.Printf("Promo creation error: %v", err)
	}
}

// UsePromo Function to use the Promo Code
func UsePromo(update *tgbotapi.Update, promo string) bool {
	foundPromo, err := repositories.GetPromo(promo)
	if err != nil {
		return false // Promo not found
	}
	if int(foundPromo.UseCount)-foundPromo.Activations == 0 {
		return false
	}
	user, err := repositories.GetUserByTgID(update.SentFrom().ID)
	if err != nil {
		return false // User not found
	}
	usedPromo := models.UsedPromo{
		PromoID: foundPromo.ID,
		UserID:  user.ID,
	}
	err = repositories.CreateUsedPromo(&usedPromo)
	if err != nil {
		return false // It was already activated
	}
	ChangeBalance(foundPromo.Amount, update)
	foundPromo.UseCount++
	repositories.UpdatePromo(&foundPromo)
	return true
}
