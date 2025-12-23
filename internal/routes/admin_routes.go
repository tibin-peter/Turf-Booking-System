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
		//logout and dashboar
		adminPanel.GET("/logout", adminH.AdminLogout)
		adminPanel.GET("/dashboard", adminH.ShowDashboardPage)

		//turf related
		adminPanel.GET("/turfs", adminH.AdminShowTurfs)
		adminPanel.GET("/turfs/add", adminH.AdminShowAddTurfPage)
		adminPanel.POST("/turfs/add", adminH.AdminAddTurf)
		adminPanel.GET("/turfs/edit/:id", adminH.AdminShowEditTurfPage)
		adminPanel.POST("/turfs/edit/:id", adminH.AdminEditTurf)
		adminPanel.GET("/turfs/delete/:id", adminH.AdminDeleteTurf)

		//slot related
		adminPanel.GET("/turfs/:id/slots", adminH.ListSlots)
		adminPanel.GET("/turfs/:id/slots/filter", adminH.FilterSlotsByDate)
		adminPanel.POST("/turfs/:id/slots", adminH.AddNewSlot)
		adminPanel.GET("/slots/:id/edit", adminH.ShowEditSlotPage)
		adminPanel.POST("/slots/:id/edit", adminH.EditSlot)
		adminPanel.GET("/slots/:id/delete", adminH.DeleteSlot)

		//booking related
		adminPanel.GET("/bookings", adminH.ListBookings)
		adminPanel.GET("/bookings/:id/approve", adminH.ApproveBooking)
		adminPanel.GET("/bookings/:id/cancel", adminH.CancelBooking)

		//payment related
		adminPanel.GET("/payments", adminH.ListPayments)
		adminPanel.GET("/payments/:id/approve", adminH.ApprovePayment)

		//user related
		adminPanel.GET("/users", adminH.ListUsers)
		adminPanel.GET("users/:user_id/block", adminH.BlockUser)
		adminPanel.GET("users/:user_id/unblock", adminH.UnblockUser)
	}
}
