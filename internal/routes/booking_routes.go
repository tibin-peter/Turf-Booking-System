package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

func BookingRoutes(r *gin.Engine, bookingH *handlers.BookingHandler, repo repository.Repository) {
	bookings := r.Group("/bookings")
	bookings.Use(middleware.AuthRequired(repo))
	{
		bookings.POST("/", bookingH.CreateBooking)
		bookings.GET("/my", bookingH.ListBookings)
		bookings.DELETE("/:id", bookingH.CancelBooking)
		bookings.POST("/:id/pay", bookingH.ConfirmPayment)
	}
}
