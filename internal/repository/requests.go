package repositories

import (
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
)

// CREATE

// CreateRequest Creates new request
func CreateRequest(request *models.FreeRequest) (err error) {
	err = DB.Create(request).Error
	return
}
