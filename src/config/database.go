package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectionDB() {
	db, err := sql.Open("mysql", "root:root@/go_products?parseTime=true")
	if err != nil {
	panic(err);
	}

	log.Println("Database connected")
	DB = db
}