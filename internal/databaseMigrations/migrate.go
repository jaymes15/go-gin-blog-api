package databasemigrations

import (
	"blog/internal/models"
	"blog/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&models.User{}, &models.Article{})

	if err != nil {
		log.Fatalf(":::::Error Reading Configs::::: %s", err)
	}

	fmt.Println("Migration done")
}
