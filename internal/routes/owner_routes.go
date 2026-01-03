package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/middleware"
	"github.com/tibin-peter/Turf-Booking-System/internal/owners"
)

func RegisterOwnerRoutes(r *gin.Engine, ownerH *owners.OwnerHandler) {

	// LOGIN (No Middleware)
	r.GET("/owner/login", ownerH.ShowLoginPage)
	r.POST("/owner/login", ownerH.OwnerLogin)

	// PROTECTED ADMIN ROUTES
	ownerPanel := r.Group("/owner")
	ownerPanel.Use(middleware.OwnerAuthMiddleware())
	{
		//logout and dashboar
		ownerPanel.GET("/logout", ownerH.OwnerLogout)
		ownerPanel.GET("/dashboard", ownerH.OwnerShowDashboard)

		//turf related
		ownerPanel.GET("/turfs", ownerH.ListMyTurfs)

		//slot related
		ownerPanel.GET("/turfs/:id/slots", ownerH.ListSlots)
		ownerPanel.POST("/turfs/:id/slots", ownerH.AddSlot)

		ownerPanel.GET("/slots/:id/edit", ownerH.ShowEditSlot)
		ownerPanel.POST("/slots/:id/edit", ownerH.EditSlot)

		ownerPanel.GET("/slots/:id/delete", ownerH.DeleteSlot)
		//booking related
		ownerPanel.GET("/bookings", ownerH.ListBookings)
		ownerPanel.GET("/bookings/:id/approve", ownerH.ApproveBooking)
		ownerPanel.GET("/bookings/:id/cancel", ownerH.CancelBooking)

		//payment related
		ownerPanel.GET("/payments", ownerH.ListPayments)
		ownerPanel.GET("/payments/:id/approve", ownerH.ApprovePayment)

	}
}
