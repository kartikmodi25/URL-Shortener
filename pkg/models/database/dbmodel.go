package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalURL string `json:"originalURL"`
	ShortenedURL string `json:"shortenedURL"`
}