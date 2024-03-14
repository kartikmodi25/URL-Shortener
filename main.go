package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/kartikmodi25/URL-Shortener/internal/api"
	pg "github.com/kartikmodi25/URL-Shortener/pkg/database/postgres"
)

const DSN = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata"
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	db, err := pg.GetConnection()
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
		return
	}
	err = pg.AutoMigrate(db)
	if err != nil {
		log.Fatal("failed to create tables in database", err)
		return
	}

	router := gin.Default()
	router.POST("/shorturl", api.ShortenURL(db))
	router.GET("/ping", ping)
	router.GET("/:id", api.Redirect(db))


	portStr := "8080"
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatal("invalid port number", err)
		return
	}
	if err := router.Run(":" + strconv.Itoa(port)); err != nil {
		return
	}
}
