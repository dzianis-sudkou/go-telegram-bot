package services

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddNewPromo(code string, amount int) {
	var newPromo = models.Promo{
		Code:     code,
		Amount:   amount,
		UseCount: 0,
	}

	err := repositories.CreatePromo(&newPromo)
	if err != nil {
		log.Printf("Promo creation error: %v", err)
	}
}

func UsePromo(update *tgbotapi.Update, promo string) bool {
	foundPromo, err := repositories.GetPromo(promo)
	if err != nil {
		return false // Promo not found
	}
	user, err := repositories.GetUserByTgId(update.SentFrom().ID)
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
	foundPromo.UseCount += 1
	repositories.UpdatePromo(&foundPromo)
	return true
}
