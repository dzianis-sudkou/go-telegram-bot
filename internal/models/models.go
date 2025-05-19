package models

import "time"

// 1. Table user in database
type User struct {
	ID                   uint `gorm:"primaryKey"`
	ChatId               int64
	TgId                 int64
	FullName             string
	MsgCount             uint
	FreeRequestCount     uint
	Credits              int
	GeneratedImagesCount uint
	RegistrationDate     time.Time
	State                string
	Authorized           bool
	FreeRequests         []FreeRequest
}

// 2. Table posts in database
type Post struct {
	ID          uint `gorm:"primaryKey"`
	Description string
	Images      []Image
}

// 3 Table images in database
type Image struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	ImageID string
	PostID  uint
}

// 4. Table free_requests in database
type FreeRequest struct {
	ID           uint `gorm:"primaryKey"`
	Text         string
	CreationDate time.Time
	Language     string
	UserID       uint
}

// 5. Table generated_images in database
type GeneratedImage struct {
	ID       uint `gorm:"primaryKey"`
	Message  int64
	Prompt   string
	TaskUUID string
	ImageURL string
	Done     bool
	Chat     int64
	Model    string
	Language string
	NSFW     bool
}

// 6. RU locales table
type RuLocale struct {
	ID    uint   `gorm:"primaryKey"`
	State string `gorm:"uniqueIndex"`
	Text  string
}

// 7. EN locales table
type EnLocale struct {
	ID    uint   `gorm:"primaryKey"`
	State string `gorm:"uniqueIndex"`
	Text  string
}

type Payment struct {
	ID                      uint `gorm:"primaryKey"`
	Currency                string
	TotalAmount             int
	InvoicePayload          string
	ShippingOptionID        string
	TelegramPaymentChargeId string
	ProviderPaymentChargeId string
}
