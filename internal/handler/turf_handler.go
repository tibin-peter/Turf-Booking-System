package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func GetAllTurfs(c *gin.Context) {
	turf, err := service.ListTurfs()

	if err != nil {
		utils.JSONError(c, 500, err.Error())
		return
	}
	utils.JSONSuccess(c, "Turf fetched successfully", turf)
}

func GetTurfByID(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam)

	if err != nil {
		utils.JSONError(c, 400, "invalid turf id")
		return
	}
	turf, err := service.GetTurfByID(uint(idInt))

	if err != nil {
		utils.JSONError(c, 400, "turf not found")
		return
	}
	utils.JSONSuccess(c, "Turf fetched successfully", turf)

}
