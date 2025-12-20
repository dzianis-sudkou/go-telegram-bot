package services

import (
	"log"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	repositories "github.com/dzianis-sudkou/go-telegram-bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// AddNewPayment Adds new payment to the database
func AddNewPayment(payment *tgbotapi.SuccessfulPayment) {
	newPayment := models.Payment{
		Currency:                payment.Currency,
		TotalAmount:             payment.TotalAmount,
		InvoicePayload:          payment.InvoicePayload,
		ShippingOptionID:        payment.ShippingOptionID,
		TelegramPaymentChargeId: payment.TelegramPaymentChargeID,
		ProviderPaymentChargeId: payment.ProviderPaymentChargeID,
	}

	if err := repositories.CreatePayment(&newPayment); err != nil {
		log.Printf("Payment create: %v", err)
	}
}
