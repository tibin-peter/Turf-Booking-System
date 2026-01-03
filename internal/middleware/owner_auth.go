package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func OwnerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			c.Redirect(http.StatusFound, "/owner/login")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil || claims.Role != "owner" {
			c.Redirect(http.StatusFound, "/owner/login")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
