package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func BookingRoutes(r *gin.Engine) {
	bookings := r.Group("/bookings")
	bookings.Use(middleware.AuthRequired())
	{
		bookings.POST("/", handlers.CreateBooking)
		bookings.GET("/my", handlers.ListBookings)
		bookings.DELETE("/:id", handlers.CancelBooking)
		bookings.POST("/:id/pay", handlers.ConfirmPayment)
	}
}
