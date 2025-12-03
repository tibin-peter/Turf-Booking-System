package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

// func for get available slot by turfid
func GetSlotsByTurfID(c *gin.Context) {
	//getting the id
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

// func for get slots by the date
func GetSlotByTurfIDAndDate(c *gin.Context) {
	//extract turf id
	turfIDParam := c.Param("turfID")
	turfID, err := strconv.Atoi(turfIDParam)
	if err != nil {
		utils.JSONError(c, 400, "invalid turf id")
		return
	}

	//geting the date from query param
	date := c.Query("date")
	if date == "" {
		utils.JSONError(c, 400, "date is required")
		return
	}
	slots, err := service.ListSlotByDate(uint(turfID), date)
	if err != nil {
		utils.JSONError(c, 400, "no slots found")
		return
	}
	utils.JSONSuccess(c, "slots fetched", gin.H{"slots": slots})
}
