package middlewares

import (
	"blog/pkg/sessions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		if userID != 0 {
			c.JSON(http.StatusFound, gin.H{"message": "OK"})
			return
		}
		// before request

		c.Next()
	}
}
