package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
)

func SlotRoutes(r *gin.Engine) {
	slots := r.Group("/slots")
	{
		slots.GET("/:turfID", handlers.GetSlotsByTurfID)
	}
}
