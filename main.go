package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" 
	"github.com/kartikmodi25/URL-Shortener/util"
)

func main(){
	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("Cannot connect to config:", err)
	}
	fmt.Println(config.DBDriver)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	fmt.Println(conn)
}