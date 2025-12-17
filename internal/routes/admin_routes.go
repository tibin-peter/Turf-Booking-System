package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/admin"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
)

func RegisterAdminRoutes(r *gin.Engine, adminH *admin.AdminHandler) {

	// LOGIN (No Middleware)
	r.GET("/admin/login", adminH.ShowLoginPage)
	r.POST("/admin/login", adminH.AdminLogin)

	// PROTECTED ADMIN ROUTES
	adminPanel := r.Group("/admin")
	adminPanel.Use(middleware.AdminAuthMiddleware())
	{
		adminPanel.GET("/logout", adminH.AdminLogout)
		adminPanel.GET("/dashboard", adminH.ShowDashboardPage)

		adminPanel.GET("/turfs", adminH.AdminShowTurfs)
		adminPanel.GET("/turfs/add", adminH.AdminShowAddTurfPage)
		adminPanel.POST("/turfs/add", adminH.AdminAddTurf)
		adminPanel.GET("/turfs/edit/:id", adminH.AdminShowEditTurfPage)
		adminPanel.POST("/turfs/edit/:id", adminH.AdminEditTurf)
		adminPanel.GET("/turfs/delete/:id", adminH.AdminDeleteTurf)
	}
}
