package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

func ShowDashboardPage(c *gin.Context) {

	totalUsers := repository.CountUsers()
	totalTurfs := repository.CountTurfs()
	totalSlots := repository.CountSlots()
	totalBookings := repository.CountBookings()

	today := time.Now().Format("2000-01-01")
	todayBookings := repository.CountBookingsByDate(today)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"TotalUsers":    totalUsers,
		"TotalTurfs":    totalTurfs,
		"TotalSlots":    totalSlots,
		"TotalBookings": totalBookings,
		"TodayBookings": todayBookings,
	})
}
