package respositories

import "blog/internal/models"

type UserRespositoryInterface interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) models.User
	FindByID(id int) models.User
}
