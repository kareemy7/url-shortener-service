package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	UserID      uint
	OriginalURL string
	ShortCode   string
}
