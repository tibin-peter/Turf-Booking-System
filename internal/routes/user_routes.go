package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

func RegisterUserRoutes(r *gin.Engine, authH *handlers.AuthHandler, userH *handlers.UserHandler, repo repository.Repository) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authH.Register)
		auth.POST("/login", authH.Login)
		auth.GET("/refresh", authH.Refresh)
		auth.POST("/logout", authH.Logout)
	}

	protected := r.Group("/user")
	protected.Use(middleware.AuthRequired(repo))
	{
		protected.GET("/profile", userH.GetProfile)
		protected.PUT("/update", userH.UpdateProfile)
		protected.GET("/bookings", userH.BookingHistory)
	}
}
