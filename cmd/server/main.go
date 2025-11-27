package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Printf("Server running at port 8080")

	r.Run(":8080")
}
