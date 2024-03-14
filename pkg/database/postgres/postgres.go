package postgres

import (
	"fmt"
	"log"
	"math/rand"

	models "github.com/kartikmodi25/URL-Shortener/pkg/models/database"
	"github.com/kartikmodi25/URL-Shortener/util"
	"github.com/pkg/errors"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

const DSN = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata"

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.URL{})
	return errors.Wrap(err, "db.AutoMigrate")
}

func GetConnection() (*gorm.DB, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot connect to config:", err)
	}
	dsn := fmt.Sprintf(DSN, config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "db.GetConnection()")
	}
	return db, nil
}

func GenerateShortURL(db *gorm.DB, originalURL string) (string, error) {
	var url models.URL
	err := db.Model(&models.URL{}).Where(&models.URL{OriginalURL: originalURL}).First(&url).Error
	if err == nil {
		return url.ShortenedURL, nil
	}
	baseURL := "shortURL.com/"
	urlSuffix := RandStringBytes(6)
	baseURL += urlSuffix
	url.OriginalURL = originalURL
	url.ShortenedURL = baseURL
	err = db.Model(&models.URL{}).Save(&url).Error
	if err != nil {
		return "", err
	}
	return baseURL, nil
}
func FindURL(db *gorm.DB, shortenedURL string) (string, error) {
	url := models.URL{}
	err := db.Where(&models.URL{ShortenedURL: shortenedURL}).First(&url).Error
	if err != nil{
		return "", err
	}
	return url.OriginalURL, err
}