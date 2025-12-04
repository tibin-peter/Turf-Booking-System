package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	routes.RegisterUserRoutes(r)
	routes.SlotRoutes(r)
	routes.TurfRoutes(r)
	routes.BookingRoutes(r)
	routes.RegisterAdminRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(r.Run(":" + port))
}
