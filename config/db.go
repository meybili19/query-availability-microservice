package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	server := os.Getenv("DB_SERVER")
	dbName := os.Getenv("DB_NAME")

	// Ensure that user and password are not empty
	if user == "" || password == "" || server == "" || dbName == "" {
		log.Fatal("❌ Missing database credentials in environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, server, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	log.Println("✅ Successfully connected to the database")
}
