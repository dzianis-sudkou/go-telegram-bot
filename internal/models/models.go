package models

import "time"

// 1. Table user in database
type User struct {
	ID                   uint
	TgId                 uint64
	FullName             string
	MsgCount             uint
	FreeRequestCount     uint
	GeneratedImagesCount uint
	RegistrationDate     time.Time
	State                string
	Authorized           bool
	FreeRequests         []FreeRequest
	GeneratedImages      []GeneratedImage
}

// 2. Table posts in database
type Post struct {
	ID          uint
	Description string
	Images      []Image
}

// 3 Table images in database
type Image struct {
	ID      uint
	Name    string
	ImageID string
	PostID  uint
}

// 4. Table free_requests in database
type FreeRequest struct {
	ID           uint
	Text         string
	CreationData time.Time
	Language     string
	UserID       uint
}

// 5. Table generated_images in database
type GeneratedImage struct {
	ID       uint
	TaskUUID string
	ImageURL string
	Done     bool
	UserID   uint
}
