package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/routes"
)

func main() {

	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, "..", "..", ".env")
	godotenv.Load(envPath)
	config.ConnectDB()

	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routes.RegisterUserRoutes(r)

	log.Fatal(r.Run(":" + port))
}
