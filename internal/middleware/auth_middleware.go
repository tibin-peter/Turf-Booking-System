package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func AuthRequired(repo repository.Repository) gin.HandlerFunc {
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

		var user model.User
		if err := repo.FindById(&user, claims.UserID); err != nil {
			utils.JSONError(c, http.StatusUnauthorized, "user not found")
			c.Abort()
			return
		}

		// 4. Block check
		if user.IsBlocked {
			utils.JSONError(c, http.StatusForbidden, "account blocked by admin")
			c.Abort()
			return
		}

		c.Set("user_id", user.ID)
		c.Set("role", user.Role)

		c.Next()
	}
}
