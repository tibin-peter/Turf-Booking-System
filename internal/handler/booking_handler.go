package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

// func for create booking
func CreateBooking(c *gin.Context) {
	uid, exist := c.Get("user_id")
	if !exist {
		utils.JSONError(c, 401, "unauthorized: user id missing")
		return
	}
	userID := uid.(uint)
	//binding the body
	var body struct {
		TurfID        uint   `json:"turf_id"`
		SlotID        uint   `json:"slot_id"`
		Amount        int    `json:"amount"`
		PaymentMethod string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.JSONError(c, 400, "invalid booking data")
		return
	}

	//create booking object

	booking := model.Booking{
		UserID:        userID,
		TurfID:        body.TurfID,
		SlotID:        body.SlotID,
		Amount:        body.Amount,
		PaymentMethod: body.PaymentMethod,
	}
	//call service for logic

	if err := service.CreteUserBooking(&booking); err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}
	//response

	utils.JSONSuccess(c, "booking created successfully", gin.H{
		"booking": booking,
	})
}

// list user bookings
func ListBookings(c *gin.Context) {
	uid, exist := c.Get("user_id")
	if !exist {
		utils.JSONError(c, 401, "unauthorized user")
		return
	}
	userID := uid.(uint)

	bookings, err := service.ListUserBookings(userID)
	if err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}

	utils.JSONSuccess(c, "bookings fetched", gin.H{
		"bookings": bookings,
	})
}

// func for cancel booking
func CancelBooking(c *gin.Context) {
	uid, exists := c.Get("user_id")
	if !exists {
		utils.JSONError(c, 401, "unauthorized")
		return
	}
	userID := uid.(uint)

	//read booking id from url
	idParam := c.Param("id")
	bid, err := strconv.Atoi(idParam)
	if err != nil {
		utils.JSONError(c, 400, "invalid booking id")
		return
	}
	//calling service
	if err := service.CancelUserBooking(uint(bid), userID); err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}
	utils.JSONSuccess(c, "booking cancelled successfully", nil)
}
