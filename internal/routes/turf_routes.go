package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/handler"
)

func TurfRoutes(r *gin.Engine) {
	turf := r.Group("/turfs")
	{
		turf.GET("/", handlers.GetAllTurfs)
		turf.GET("/:id", handlers.GetTurfByID)
	}
}
