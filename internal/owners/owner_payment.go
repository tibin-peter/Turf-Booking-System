package owners

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// List payments for owner turfs only
func (h *OwnerHandler) ListPayments(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	if ownerID == 0 {
		c.Redirect(http.StatusFound, "/owner/login")
		return
	}

	var payments []model.Payment

	err := h.repo.FindMany(
		&payments,
		"booking_id IN (SELECT id FROM bookings WHERE turf_id IN (SELECT id FROM turfs WHERE owner_id = ?))",
		ownerID,
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "owner_payments.html", gin.H{
			"error": "failed to load payments",
		})
		return
	}

	c.HTML(http.StatusOK, "owner_payments.html", gin.H{
		"Payments": payments,
	})
}

// Approve payment (
func (h *OwnerHandler) ApprovePayment(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var payment model.Payment
	if err := h.repo.FindById(&payment, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/owner/payments")
		return
	}

	var booking model.Booking
	if err := h.repo.FindById(&booking, payment.BookingID); err != nil {
		c.Redirect(http.StatusFound, "/owner/payments")
		return
	}

	var turf model.Turf
	if err := h.repo.FindById(&turf, booking.TurfID); err != nil || turf.OwnerID != ownerID {
		c.Redirect(http.StatusFound, "/owner/payments")
		return
	}

	if payment.Status == "pending" {
		payment.Status = "paid"
		_ = h.repo.Update(&payment)

		booking.PaymentStatus = "paid"
		booking.Status = "confirmed"
		_ = h.repo.Update(&booking)
	}

	c.Redirect(http.StatusFound, "/owner/payments")
}
