package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}
