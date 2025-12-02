package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routes.RegisterUserRoutes(r)
	routes.TurfRoutes(r)
	routes.SlotRoutes(r)

	log.Fatal(r.Run(":" + port))
}
