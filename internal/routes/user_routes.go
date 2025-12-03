package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func RegisterUserRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/refresh", handlers.Refresh)
		auth.POST("/logout", handlers.Logout)
	}

	protected := r.Group("/user")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/profile", handlers.GetProfile)
		protected.PUT("/update", handlers.UpdateProfile)
		protected.GET("/bookings", handlers.BookingHistory)
	}
}
