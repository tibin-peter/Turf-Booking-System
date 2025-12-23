package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *AdminHandler) ListPayments(c *gin.Context) {

	var payments []model.Payment
	if err := h.repo.FindMany(&payments, "1 = 1"); err != nil {
		c.HTML(500, "payments.html", gin.H{"error": "failed to load"})
		return
	}

	c.HTML(200, "payments.html", gin.H{
		"Payments": payments,
	})
}

//func for approve the payment by the admin
func (h *AdminHandler) ApprovePayment(c *gin.Context) {
	idStr := c.Param("id")
	paymentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 1. Fetch payment
	var payment model.Payment
	if err := h.repo.FindById(&payment, uint(paymentID)); err != nil {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 2. Prevent double approval
	if payment.Status == "paid" {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 3. Fetch booking
	var booking model.Booking
	if err := h.repo.FindById(&booking, payment.BookingID); err != nil {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 4. Validate booking state
	if booking.Status == "cancelled" {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	if booking.TotalAmount <= 0 {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 5. Approve payment
	payment.Amount = booking.TotalAmount
	payment.Status = "paid"

	if err := h.repo.Update(&payment); err != nil {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	// 6. Update booking
	booking.PaymentStatus = "paid"
	booking.Status = "confirmed"

	if err := h.repo.Update(&booking); err != nil {
		c.Redirect(http.StatusFound, "/admin/payments")
		return
	}

	c.Redirect(http.StatusFound, "/admin/payments")
}
