package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
)

func TurfRoutes(r *gin.Engine, turfH *handlers.TurfHandler) {
	turf := r.Group("/turfs")
	{
		turf.GET("/", turfH.GetAllTurfs)
		turf.GET("/:id", turfH.GetTurfByID)
	}
}
