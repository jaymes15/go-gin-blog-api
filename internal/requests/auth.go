package requests

import "mime/multipart"

// Binding from JSON
type Register struct {
	Name     string                `form:"name" json:"name" binding:"required,min=2"`
	Email    string                `form:"email" json:"email" binding:"required,email"`
	Password string                `form:"password" json:"password"  binding:"required,min=5"`
	Image    *multipart.FileHeader `form:"image" json:"image"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required,min=5"`
}
