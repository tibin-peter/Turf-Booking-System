package main

import (
	// "log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tibin-peter/Turf-Booking-System/config"
)

func main() {

	godotenv.Load()
	config.ConnectDB()

	// r := routes.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// log.Fatal(r.Run(":" + port))
}
