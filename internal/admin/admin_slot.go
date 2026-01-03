package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// func for list slots for a particular turf
func (h *AdminHandler) ListSlots(c *gin.Context) {
	idStr := c.Param("id")
	turfid, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid turf id",
		})
		return
	}

	var slots []model.TimeSlot
	if err := h.repo.FindMany(&slots, "turf_id = ?", uint(turfid)); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_slots.html", gin.H{
			"error": "failed to load slots",
		})
		return
	}
	c.HTML(http.StatusOK, "admin_slots.html", gin.H{
		"Slots":  slots,
		"TurfID": turfid,
	})
}

// filter slots by date
func (h *AdminHandler) FilterSlotsByDate(c *gin.Context) {
	idStr := c.Param("id")
	turfid, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid turf id",
		})
		return
	}

	date := c.Query("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid date format",
		})
		return
	}

	var slots []model.TimeSlot
	if err := h.repo.FindMany(
		&slots,
		"turf_id = ? AND day = ?",
		uint(turfid),
		parsedDate,
	); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_slots.html", gin.H{
			"error": "failed to fetch the slots",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_slots.html", gin.H{
		"Slots":  slots,
		"TurfID": turfid,
		"Date":   date,
	})
}

// function for add a new slot
func (h *AdminHandler) AddNewSlot(c *gin.Context) {
	//get turf id from url
	idstr := c.Param("id")
	turfid, err := strconv.Atoi(idstr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid turf id",
		})
		return
	}
	//get form values
	daystr := c.PostForm("day")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	//validate the field
	var slots []model.TimeSlot
	sloterr := h.repo.FindMany(&slots, "turf_id = ?", uint(turfid))
	if sloterr != nil || daystr == "" || startTime == "" || endTime == "" {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "all fields are required",
		})
		return
	}
	//parse string into time.time
	day, err := time.Parse("2006-01-02", daystr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error":  "invalid date format",
			"Slots":  slots,
			"TurfID": turfid,
		})
		return
	}
	//normalize today
	now := time.Now()
	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
	//avoid past dats creation
	if day.Before(today) {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error":  "cannot add slots for past dates",
			"Slots":  slots,
			"TurfID": turfid,
		})
		return
	}
	//parse times
	start, err1 := time.Parse("15:04", startTime)
	end, err2 := time.Parse("15:04", endTime)

	if err1 != nil || err2 != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid format",
		})
		return
	}
	if !start.Before(end) {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error":  "start time must be before end time",
			"Slots":  slots,
			"TurfID": turfid,
		})
		return
	}
	//avoid past time creation
	if day.Equal(today) {
		currentTime, _ := time.Parse("15:04", now.Format("15:04"))
		if start.Before(currentTime) {
			c.HTML(http.StatusBadGateway, "admin_slots.html", gin.H{
				"error":  "cannot add slots in the past time",
				"Slots":  slots,
				"TurfID": turfid,
			})
			return
		}
	}
	//check the slot are existing the same turf same day same time
	var count int64
	count, err = h.repo.Count(
		&model.TimeSlot{},
		`turf_id = ?
		AND day = ?
		AND (
		(start_time < ? AND end_time > ?)
		OR (start_time < ? AND end_time > ?)
		OR (start_time >= ? AND end_time <= ?)
		)
		`,
		uint(turfid),
		day,
		endTime, startTime,
		endTime, startTime,
		startTime, endTime,
	)
	if count > 0 {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error":  "already have a slot",
			"Slots":  slots,
			"TurfID": turfid,
		})
		return
	}
	//create slot
	slot := model.TimeSlot{
		TurfID:    uint(turfid),
		Day:       day,
		StartTime: startTime,
		EndTime:   endTime,
	}

	//insert into db
	if err := h.repo.Insert(&slot); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_slots.html", gin.H{
			"error":  "failed to add the slot",
			"Slots":  slots,
			"TurfID": turfid,
		})
		return
	}

	//redirect into the slots page
	c.Redirect(http.StatusFound, "/admin/turfs/"+idstr+"/slots")

}

// Show edit slot page (GET)
func (h *AdminHandler) ShowEditSlotPage(c *gin.Context) {

	slotIDStr := c.Param("id")
	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "admin_editslot.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "cannot edit a booked slot",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_editslot.html", gin.H{
		"Slot": slot,
	})
}

// function for edit slot
func (h *AdminHandler) EditSlot(c *gin.Context) {
	slotIDStr := c.Param("id")
	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "admin_editslot.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "cannot edit booked slot",
			"Slots": slot,
		})
		return
	}

	dayStr := c.PostForm("day")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	//validation
	if dayStr == "" || startTime == "" || endTime == "" {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "all fields are required",
			"Slot":  slot,
		})
		return
	}

	day, err := time.Parse("2006-01-02", dayStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_editslot.html", gin.H{
			"error": "invalid date format",
			"Slot":  slot,
		})
		return
	}

	slot.Day = day
	slot.StartTime = startTime
	slot.EndTime = endTime

	if err := h.repo.Update(&slot); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_editslot.html", gin.H{
			"error": "failed to update slot",
			"Slot":  slot,
		})
		return
	}

	c.Redirect(
		http.StatusFound,
		"/admin/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots",
	)
}

// function for delete a slot by its id
func (h *AdminHandler) DeleteSlot(c *gin.Context) {
	idStr := c.Param("id")
	slotID, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "admin_slots.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "admin_slots.html", gin.H{
			"error": "cannot delete booked slot",
		})
		return
	}

	if err := h.repo.Delete(&model.TimeSlot{}, "id = ?", uint(slotID)); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_slots.html", gin.H{
			"error": "failed to delete slot",
		})
		return
	}

	c.Redirect(
		http.StatusFound,
		"/admin/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots",
	)
}
