package main

import (
	"database/sql"
	"log"

	"github.com/kartikmodi25/URL-Shortener/pkg/models"
	"github.com/kartikmodi25/URL-Shortener/util"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot connect to config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{})

	db.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
}
