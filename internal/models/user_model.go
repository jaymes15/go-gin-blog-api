package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"varchar:191;not null"`
	Email    string `gorm:"varchar:191;unique;not null"`
	Password string `gorm:"varchar:191;not null"`
	Image    string `gorm:"text;"`
}
