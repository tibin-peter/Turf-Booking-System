package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/admin"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func RegisterAdminRoutes(r *gin.Engine) {

	//admin login no middleware
	r.GET("/admin/login", admin.ShowLoginPage)
	r.POST("/admin/login", admin.AdminLogin)

	//protected admin routes
	adminPanel := r.Group("/admin")
	adminPanel.Use(middleware.AdminAuthMiddleware())
	{
		adminPanel.GET("/logout", admin.AdminLogout)
		adminPanel.GET("/dashboard", admin.ShowDashboardPage)
	}
}
