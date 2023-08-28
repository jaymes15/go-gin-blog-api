package services

import (
	handlemedia "blog/internal/handleMedia"
	"blog/internal/models"
	"blog/internal/requests"
	"blog/internal/responses"
	"blog/internal/respositories"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRespository respositories.UserRespositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		userRespository: respositories.NewUserRespository(),
	}
}

func (userService *UserService) CreateUser(c *gin.Context, request requests.Register) (responses.User, error) {
	var response responses.User
	image := ""

	if request.Image != nil {
		fileIdentifier := fmt.Sprintf("User Image %s", request.Name)

		resp, err := handlemedia.UploadImage(c, request.Image, fileIdentifier)

		if err != nil {
			return response, err
		}
		image = resp.URL
	}

	newUser := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Image:    image,
	}
	user, err := userService.userRespository.Create(newUser)

	if err != nil {
		return response, err
	}

	if user.ID == 0 {
		return response, errors.New("error on creating user")
	}

	return responses.ToUser(user), nil
}

func (userService *UserService) LoginUser(request requests.Login) (responses.User, error) {
	var response responses.User

	user := userService.userRespository.FindByEmail(request.Email)

	if user.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return responses.ToUser(user), nil
}
