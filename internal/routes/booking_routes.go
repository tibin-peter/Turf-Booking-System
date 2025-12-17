package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func BookingRoutes(r *gin.Engine, bookingH *handlers.BoookinHandler) {
	bookings := r.Group("/bookings")
	bookings.Use(middleware.AuthRequired())
	{
		bookings.POST("/", bookingH.CreateBooking)
		bookings.GET("/my", bookingH.ListBookings)
		bookings.DELETE("/:id", bookingH.CancelBooking)
		bookings.POST("/:id/pay", bookingH.ConfirmPayment)
	}
}
