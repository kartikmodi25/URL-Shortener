package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ShortenedURL string `gorm:"not null"`
}