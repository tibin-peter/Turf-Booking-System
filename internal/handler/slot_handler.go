package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func GetSlotsByTurfID(c *gin.Context) {
	turfIDParam := c.Param("turfID")
	id, err := strconv.Atoi(turfIDParam)
	if err != nil {
		utils.JSONError(c, 400, "invalid turf id")
		return
	}
	slots, err := service.ListSlotsByTurfID(uint(id))
	if err != nil {
		utils.JSONError(c, 400, "slot not found")
		return
	}
	utils.JSONSuccess(c, "slots fetched successfully", slots)
}
