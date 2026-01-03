package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *AdminHandler) ListBookings(c *gin.Context) {

	var bookings []model.Booking
	if err := h.repo.FindMany(&bookings, "1 = 1"); err != nil {
		c.HTML(500, "admin_bookings.html", gin.H{"error": "failed to load"})
		return
	}

	c.HTML(200, "admin_bookings.html", gin.H{
		"Bookings": bookings,
	})
}
func (h *AdminHandler) ForceCancelBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var booking model.Booking
	if err := h.repo.FindById(&booking, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/admin/bookings")
		return
	}

	if booking.Status == "cancelled" {
		c.Redirect(http.StatusFound, "/admin/bookings")
		return
	}

	now := time.Now()
	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"
	booking.CancelledBy = "admin"
	booking.CancelledAt = &now

	_ = h.repo.Update(&booking)

	c.Redirect(http.StatusFound, "/admin/bookings")
}
