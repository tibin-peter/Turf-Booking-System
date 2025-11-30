package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name, value string, expiry time.Time) {
	maxAge := int(time.Until(expiry).Seconds())
	c.SetCookie(
		name,
		value,
		maxAge,
		"/",
		"",
		false,
		true,
	)
}

func ClearCookie(c *gin.Context, name string) {
	c.SetCookie(
		name,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}
