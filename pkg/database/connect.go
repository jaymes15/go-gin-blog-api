package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	err := godotenv.Load()
	err2 := godotenv.Load("../../.env")
	if err != nil && err2 != nil {
		log.Fatalf("Error loading .env file %s", err)
		log.Fatalf("Error loading .env file %s", err2)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, name, port)

	dsn := dbConfig
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err)
	}

	DB = db
}
