package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
)

type BookingHandler struct {
	service *service.BookingService
}

func NewBookingHandler(s *service.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req struct {
		TurfID        uint   `json:"turf_id"`
		SlotID        uint   `json:"slot_id"`
		PaymentMethod string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.SlotID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slot_id is required"})
		return
	}

	userID := c.GetUint("user_id")

	booking := model.Booking{
		UserID:        userID,
		TurfID:        req.TurfID,
		SlotID:        req.SlotID,
		PaymentMethod: req.PaymentMethod,
	}

	if err := h.service.CreateBooking(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "booking created successfully",
		"booking_id": booking.ID,
	})
}

func (h *BookingHandler) ConfirmPayment(c *gin.Context) {

	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.ConfirmPayment(uint(id), userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "payment initiated"})
}

func (h *BookingHandler) ListBookings(c *gin.Context) {

	userID := c.GetUint("user_id")

	bookings, err := h.service.ListUserBookings(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch"})
		return
	}

	c.JSON(200, bookings)
}

func (h *BookingHandler) CancelBooking(c *gin.Context) {

	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.CancelBooking(uint(id), userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "booking cancelled"})
}
