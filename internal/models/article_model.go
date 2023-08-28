package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191;not null"`
	Content string `gorm:"text;not null"`
	UserId  uint
	User    User
}
