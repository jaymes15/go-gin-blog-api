package routes

import (
	"blog/internal/controllers"
	"blog/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := controllers.NewArticleController()
	userController := controllers.NewAuthController()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/articles", articleController.Show)
		guestGroup.GET("/articles/:id", articleController.Details)
		guestGroup.POST("/register", userController.Register)
		guestGroup.POST("/login", userController.Login)

	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/articles", articleController.Create)
		router.POST("/logout", userController.HandleLogout)

	}

}
