package utils

import "github.com/gin-gonic/gin"

func JSONSuccess(c *gin.Context, message string, data any) {
	c.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
