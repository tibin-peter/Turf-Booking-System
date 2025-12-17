package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
)

func SlotRoutes(r *gin.Engine, slotH *handlers.SlotHandler) {
	slots := r.Group("/slots")
	{
		slots.GET("/:turfID", slotH.GetSlotsByTurfID)
		slots.GET("/:turfID/date", slotH.GetSlotByTurfIDAndDate)
	}
}
