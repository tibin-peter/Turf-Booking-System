package owners

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *OwnerHandler) OwnerShowDashboard(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	// needed counts

	totalTurfs, _ := h.repo.Count(
		&model.Turf{},
		"owner_id = ?",
		ownerID,
	)

	totalSlots, _ := h.repo.Count(
		&model.TimeSlot{},
		"turf_id IN (SELECT id FROM turfs WHERE owner_id = ?)",
		ownerID,
	)

	today := time.Now().Format("2006-01-02")

	todayBookings, _ := h.repo.Count(
		&model.Booking{},
		"turf_id IN (SELECT id FROM turfs WHERE owner_id = ?) AND DATE(created_at) = ?",
		ownerID,
		today,
	)

	// weekly bookings

	var weekLabels []string
	var weekValues []int

	for i := 6; i >= 0; i-- {
		day := time.Now().AddDate(0, 0, -i)
		weekLabels = append(weekLabels, day.Format("Mon"))

		count, _ := h.repo.Count(
			&model.Booking{},
			"turf_id IN (SELECT id FROM turfs WHERE owner_id = ?) AND DATE(created_at) = ?",
			ownerID,
			day.Format("2006-01-02"),
		)

		weekValues = append(weekValues, int(count))
	}

	// monthly revenue

	var monthLabels []string
	var monthValues []int

	year := time.Now().Year()

	for m := 1; m <= 12; m++ {
		monthLabels = append(monthLabels, time.Month(m).String())

		start := time.Date(year, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 1, 0)

		var payments []model.Payment
		_ = h.repo.FindMany(
			&payments,
			`booking_id IN (
				SELECT id FROM bookings
				WHERE turf_id IN (SELECT id FROM turfs WHERE owner_id = ?)
			) AND created_at >= ? AND created_at < ? AND status = 'paid'`,
			ownerID,
			start,
			end,
		)

		total := 0
		for _, p := range payments {
			total += p.Amount
		}

		monthValues = append(monthValues, total)
	}

	// json for charts

	weekLabelsJSON, _ := json.Marshal(weekLabels)
	weekValuesJSON, _ := json.Marshal(weekValues)
	monthLabelsJSON, _ := json.Marshal(monthLabels)
	monthValuesJSON, _ := json.Marshal(monthValues)

	// ---------------- RENDER ----------------

	c.HTML(http.StatusOK, "owner_dashboard.html", gin.H{
		"TotalTurfs":    totalTurfs,
		"TotalSlots":    totalSlots,
		"TodayBookings": todayBookings,

		"WeekLabels": template.JS(weekLabelsJSON),
		"WeekValues": template.JS(weekValuesJSON),

		"MonthLabels": template.JS(monthLabelsJSON),
		"MonthValues": template.JS(monthValuesJSON),
	})
}
