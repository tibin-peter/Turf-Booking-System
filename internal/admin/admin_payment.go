package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *AdminHandler) ListPayments(c *gin.Context) {

	var payments []model.Payment
	if err := h.repo.FindMany(&payments, "1 = 1"); err != nil {
		c.HTML(500, "admin_payments.html", gin.H{"error": "failed to load"})
		return
	}

	c.HTML(200, "admin_payments.html", gin.H{
		"Payments": payments,
	})
}
