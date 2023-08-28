package services

import (
	"blog/internal/requests"
	"blog/internal/responses"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	CreateUser(c *gin.Context, request requests.Register) (responses.User, error)
	LoginUser(request requests.Login) (responses.User, error)
}
