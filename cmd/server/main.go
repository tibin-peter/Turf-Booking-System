package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/admin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/routes"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
)

func main() {

	// 1. Connect DB
	config.ConnectDB()

	// 2. Create gin instance
	r := gin.Default()

	// 3. Load templates (admin SSR)
	r.SetHTMLTemplate(loadTemplates())

	// 4. Create repository instance
	repo := repository.Newrepo(config.DB)

	// 5. Create service instances
	authService := service.NewAuthService(repo)
	userService := service.NewUserService(repo)
	turfService := service.NewTurfService(repo)
	slotService := service.NewSlotService(repo)
	bookingService := service.NewBookingService(repo)

	// 6. Create handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	turfHandler := handlers.NewTurfHandler(turfService)
	slotHandler := handlers.NewSlotHandler(slotService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	adminHandler := admin.NewAdminHandler(repo)

	// 7. Register routes
	routes.RegisterUserRoutes(r, authHandler, userHandler, repo)
	routes.TurfRoutes(r, turfHandler)
	routes.SlotRoutes(r, slotHandler)
	routes.BookingRoutes(r, bookingHandler, repo)
	routes.RegisterAdminRoutes(r, adminHandler)

	// 8. Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(r.Run(":" + port))
}

func loadTemplates() *template.Template {
	tmpl := template.New("")
	template.Must(tmpl.ParseGlob("templates/*.html"))
	return tmpl
}
