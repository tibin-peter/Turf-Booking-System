package admin

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// weelky bookings line chart

func (h *AdminHandler) getWeeklyBookingsData() ([]string, []int, error) {
	var labels []string
	var values []int

	for i := 6; i >= 0; i-- {
		date := time.Now().AddDate(0, 0, -i)
		labels = append(labels, date.Format("Mon"))

		count, err := h.repo.Count(
			&model.Booking{},
			"DATE(created_at) = ?",
			date.Format("2006-01-02"),
		)
		if err != nil {
			return nil, nil, err
		}

		values = append(values, int(count))
	}

	return labels, values, nil
}

// turf usage pie chart

func (h *AdminHandler) getTurfUsageStats() ([]string, []int, error) {
	var labels []string
	var values []int

	var turfs []model.Turf
	if err := h.repo.FindMany(&turfs, "1 = 1"); err != nil {
		return nil, nil, err
	}

	for _, t := range turfs {
		count, err := h.repo.Count(
			&model.Booking{},
			"turf_id = ?",
			t.ID,
		)
		if err != nil {
			return nil, nil, err
		}

		labels = append(labels, t.Name)
		values = append(values, int(count))
	}

	return labels, values, nil
}

// monthly revenue bar chart

func (h *AdminHandler) getMonthlyRevenueStats() ([]string, []int, error) {
	var labels []string
	var values []int

	year := time.Now().Year()

	for m := 1; m <= 12; m++ {
		labels = append(labels, time.Month(m).String())

		start := time.Date(year, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 1, 0)

		var payments []model.Payment
		if err := h.repo.FindMany(
			&payments,
			"created_at >= ? AND created_at < ?",
			start,
			end,
		); err != nil {
			return nil, nil, err
		}

		total := 0
		for _, p := range payments {
			total += p.Amount
		}

		values = append(values, total)
	}

	return labels, values, nil
}

// function for show the dashboard page

func (h *AdminHandler) ShowDashboardPage(c *gin.Context) {

	// summary count
	totalUsers, _ := h.repo.Count(&model.User{}, "")
	totalTurfs, _ := h.repo.Count(&model.Turf{}, "")
	totalSlots, _ := h.repo.Count(&model.TimeSlot{}, "")
	totalBookings, _ := h.repo.Count(&model.Booking{}, "")

	today := time.Now().Format("2006-01-02")
	todayBookings, _ := h.repo.Count(
		&model.Booking{},
		"DATE(created_at) = ?",
		today,
	)

	//  chart data
	weeklyLabels, weeklyValues, _ := h.getWeeklyBookingsData()
	turfLabels, turfValues, _ := h.getTurfUsageStats()
	monthLabels, monthValues, _ := h.getMonthlyRevenueStats()

	// json marshal
	weeklyLabelsJSON, _ := json.Marshal(weeklyLabels)
	weeklyValuesJSON, _ := json.Marshal(weeklyValues)

	turfLabelsJSON, _ := json.Marshal(turfLabels)
	turfValuesJSON, _ := json.Marshal(turfValues)

	monthLabelsJSON, _ := json.Marshal(monthLabels)
	monthValuesJSON, _ := json.Marshal(monthValues)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"TotalUsers":    totalUsers,
		"TotalTurfs":    totalTurfs,
		"TotalSlots":    totalSlots,
		"TotalBookings": totalBookings,
		"TodayBookings": todayBookings,

		"WeeklyLabels": template.JS(weeklyLabelsJSON),
		"WeeklyValues": template.JS(weeklyValuesJSON),

		"TurfLabels": template.JS(turfLabelsJSON),
		"TurfValues": template.JS(turfValuesJSON),

		"MonthLabels": template.JS(monthLabelsJSON),
		"MonthValues": template.JS(monthValuesJSON),
	})
}
