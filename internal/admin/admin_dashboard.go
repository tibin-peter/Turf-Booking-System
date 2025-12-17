package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *AdminHandler) ShowDashboardPage(c *gin.Context) {

	totalUsers, _ := h.repo.Count(&model.User{}, "")
	totalTurfs, _ := h.repo.Count(&model.Turf{}, "")
	totalSlots, _ := h.repo.Count(&model.TimeSlot{}, "")
	totalBookings, _ := h.repo.Count(&model.Booking{}, "")

	today := time.Now().Format("2000-01-01")
	todayBookings, _ := h.repo.Count(&model.Booking{}, today)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"TotalUsers":    totalUsers,
		"TotalTurfs":    totalTurfs,
		"TotalSlots":    totalSlots,
		"TotalBookings": totalBookings,
		"TodayBookings": todayBookings,
	})
}
