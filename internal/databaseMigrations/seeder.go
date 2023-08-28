package databasemigrations

import (
	"blog/internal/models"
	"blog/pkg/database"
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()
	password := "test123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		log.Fatalf("unable to decode into struct %s", err)
	}
	user := models.User{Name: "Jinzhu", Email: "test@mail.com", Password: string(hashedPassword)}

	db.Create(&user)

	log.Printf("User created successfully with email address %s", user.Email)

	for i := 1; i <= 10; i++ {
		title := fmt.Sprintf("new article title %s", strconv.Itoa(i))
		content := fmt.Sprintf("new article content %s", strconv.Itoa(i))
		article := models.Article{Title: title, Content: content, UserId: user.ID}
		db.Create(&article)

		log.Printf("Article created successfully with title %s", article.Title)
	}

}
