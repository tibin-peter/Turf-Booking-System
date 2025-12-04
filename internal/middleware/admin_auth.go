package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Checking admin session...")
		session, err := c.Cookie("admin_session")
		fmt.Println("COOKIE VALUE:", session, "ERR:", err)

		if err != nil {
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
