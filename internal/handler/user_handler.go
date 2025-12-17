package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// func for get profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(uint)

	user, err := h.service.GetUserProfile(userID)
	if err != nil {
		utils.JSONError(c, 400, "user not found")
		return
	}

	user.Password = ""
	utils.JSONSuccess(c, "profile fetched", user)
}

// func for update profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(uint)

	var body model.User
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.JSONError(c, 400, "invalid input")
		return
	}

	if err := h.service.UpdateUserProfile(userID, body); err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}
	utils.JSONSuccess(c, "profile updated", nil)
}

// func for get the book hisroty
func (h *UserHandler) BookingHistory(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(uint)

	bookings, err := h.service.GetBookingHistory(userID)
	if err != nil {
		utils.JSONError(c, 400, "no hisroty found")
		return
	}
	utils.JSONSuccess(c, "history fetched", bookings)
}
