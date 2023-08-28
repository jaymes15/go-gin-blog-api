package respositories

import (
	"blog/internal/models"
	"blog/pkg/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRespository struct {
	DB *gorm.DB
}

func NewUserRespository() *UserRespository {
	return &UserRespository{
		DB: database.Connection(),
	}

}

func (userRespository *UserRespository) Create(user models.User) (models.User, error) {
	var newUser models.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return newUser, errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	userRespository.DB.Create(&user).Scan(&newUser)

	return newUser, nil
}

func (userRepository *UserRespository) FindByEmail(email string) models.User {
	var user models.User

	userRepository.DB.First(&user, "email = ?", email)

	return user
}

func (userRepository *UserRespository) FindByID(id int) models.User {
	var user models.User

	userRepository.DB.First(&user, "id = ?", id)

	return user
}
