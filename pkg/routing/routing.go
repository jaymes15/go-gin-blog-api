package routing

import (
	"blog/internal/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes(router *gin.Engine) {
	routes.Routes(router)
}
