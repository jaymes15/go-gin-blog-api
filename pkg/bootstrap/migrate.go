package bootstrap

import (
	databasemigrations "blog/internal/databaseMigrations"
	"log"

	"github.com/joho/godotenv"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databasemigrations.Migrate()

}
