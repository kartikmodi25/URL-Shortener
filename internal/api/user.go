package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartikmodi25/URL-Shortener/pkg/database/postgres"
	"github.com/kartikmodi25/URL-Shortener/pkg/models/request"
	"gorm.io/gorm"
)

func ShortenURL(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		data := request.RequestURL{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
			return
		}

		url, err := postgres.GenerateShortURL(db, data.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create shorter the URL, try again"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ShortenedURL": url})
	}
}

func Redirect(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		value := c.Param("id")
		shortenedURL := "localhost:8080/" + value
		url, err := postgres.FindURL(db, shortenedURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found!!!"})
			return
		}
		err = postgres.UpdateCount(db, shortenedURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to updated count"})
			return
		}
		c.Redirect(http.StatusFound, url)
	}
}


func TotalCount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		data := request.RequestURL{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
			return
		}

		cnt, err := postgres.GetCount(db, data.URL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found!!!"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"Count": cnt})
	}
}