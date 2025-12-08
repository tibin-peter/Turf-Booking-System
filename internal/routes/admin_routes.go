package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/admin"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func RegisterAdminRoutes(r *gin.Engine) {

	// LOGIN (No Middleware)
	r.GET("/admin/login", admin.ShowLoginPage)
	r.POST("/admin/login", admin.AdminLogin)

	// PROTECTED ADMIN ROUTES
	adminPanel := r.Group("/admin")
	adminPanel.Use(middleware.AdminAuthMiddleware())
	{
		adminPanel.GET("/logout", admin.AdminLogout)
		adminPanel.GET("/dashboard", admin.ShowDashboardPage)

		adminPanel.GET("/turfs", admin.AdminShowTurfs)
		adminPanel.GET("/turfs/add", admin.AdminShowAddTurfPage)
		adminPanel.POST("/turfs/add", admin.AdminAddTurf)
		adminPanel.GET("/turfs/edit/:id", admin.AdminShowEditTurfPage)
		adminPanel.POST("/turfs/edit/:id", admin.AdminEditTurf)
		adminPanel.GET("/turfs/delete/:id", admin.AdminDeleteTurf)
	}
}
