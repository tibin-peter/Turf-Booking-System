package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *AdminHandler) ListBookings(c *gin.Context) {

	var bookings []model.Booking
	if err := h.repo.FindMany(&bookings, "1 = 1"); err != nil {
		c.HTML(500, "bookings.html", gin.H{"error": "failed to load"})
		return
	}

	c.HTML(200, "bookings.html", gin.H{
		"Bookings": bookings,
	})
}

func (h *AdminHandler) ApproveBooking(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var booking model.Booking
	if err := h.repo.FindById(&booking, uint(id)); err != nil {
		c.Redirect(302, "/admin/bookings")
		return
	}

	if booking.Status != "approved" {
		booking.Status = "approved"
		_ = h.repo.Update(&booking)
	}

	c.Redirect(302, "/admin/bookings")
}

func (h *AdminHandler) CancelBooking(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var booking model.Booking
	if err := h.repo.FindById(&booking, uint(id)); err != nil {
		c.Redirect(302, "/admin/bookings")
		return
	}

	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"
	_ = h.repo.Update(&booking)

	c.Redirect(302, "/admin/bookings")
}
