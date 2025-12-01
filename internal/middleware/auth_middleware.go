package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			utils.JSONError(c, 401, "missing access_token")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
