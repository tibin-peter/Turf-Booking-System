package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
)

func RegisterUserRoutes(r *gin.Engine) {
	authService := service.AuthService{}
	authHandler := handler.AuthHandler{Service: authService}

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}
}
