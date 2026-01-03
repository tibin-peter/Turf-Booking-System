package owners

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// List booking owners turf only
func (h *OwnerHandler) ListBookings(c *gin.Context) {

	ownerID := c.GetUint("user_id")

	var bookings []model.Booking

	err := h.repo.FindMany(
		&bookings,
		`turf_id IN (
			SELECT id FROM turfs WHERE owner_id = ?
		)`,
		ownerID,
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "owner_bookings.html", gin.H{
			"error": "failed to load bookings",
		})
		return
	}

	c.HTML(http.StatusOK, "owner_bookings.html", gin.H{
		"Bookings": bookings,
	})
}

// Approve booking
func (h *OwnerHandler) ApproveBooking(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var booking model.Booking
	if err := h.repo.FindById(&booking, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/owner/bookings")
		return
	}

	// ownership validation
	var turf model.Turf
	if err := h.repo.FindById(&turf, booking.TurfID); err != nil || turf.OwnerID != ownerID {
		c.Redirect(http.StatusFound, "/owner/bookings")
		return
	}

	if booking.Status == "pending" {
		booking.Status = "confirmed"
		booking.PaymentStatus = "paid"
		_ = h.repo.Update(&booking)
	}

	c.Redirect(http.StatusFound, "/owner/bookings")
}

// Cancel booking
func (h *OwnerHandler) CancelBooking(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var booking model.Booking
	if err := h.repo.FindById(&booking, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/owner/bookings")
		return
	}

	// ownership validation
	var turf model.Turf
	if err := h.repo.FindById(&turf, booking.TurfID); err != nil || turf.OwnerID != ownerID {
		c.Redirect(http.StatusFound, "/owner/bookings")
		return
	}

	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"
	_ = h.repo.Update(&booking)

	//  free slot
	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, booking.SlotID); err == nil {
		slot.IsAvailable = true
		_ = h.repo.Update(&slot)
	}

	c.Redirect(http.StatusFound, "/owner/bookings")
}
