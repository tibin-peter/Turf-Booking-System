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
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "invalid turf id",
		})
		return
	}

	var slots []model.TimeSlot
	if err := h.repo.FindMany(&slots, "turf_id = ?", uint(turfid)); err != nil {
		c.HTML(http.StatusInternalServerError, "slots.html", gin.H{
			"error": "failed to load slots",
		})
		return
	}
	c.HTML(http.StatusOK, "slots.html", gin.H{
		"Slots":  slots,
		"TurfID": turfid,
	})
}

// filter slots by date
func (h *AdminHandler) FilterSlotsByDate(c *gin.Context) {
	idStr := c.Param("id")
	turfid, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
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
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
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
		c.HTML(http.StatusInternalServerError, "slots.html", gin.H{
			"error": "failed to fetch the slots",
		})
		return
	}

	c.HTML(http.StatusOK, "slots.html", gin.H{
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
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "invalid turf id",
		})
		return
	}
	//get form values
	daystr := c.PostForm("day")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	//validate the field
	if daystr == "" || startTime == "" || endTime == "" {
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "all fields are required",
		})
		return
	}
	//parse string into time.time
	day, err := time.Parse("2006-01-02", daystr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "invalid date format",
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
		c.HTML(http.StatusInternalServerError, "slots.html", gin.H{
			"error": "failed to add the slot",
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
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "edit_slot.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "cannot edit a booked slot",
		})
		return
	}

	c.HTML(http.StatusOK, "edit_slot.html", gin.H{
		"Slot": slot,
	})
}

// function for edit slot
func (h *AdminHandler) EditSlot(c *gin.Context) {
	slotIDStr := c.Param("id")
	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "edit_slot.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "cannot edit booked slot",
		})
		return
	}

	dayStr := c.PostForm("day")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	//validation
	if dayStr == "" || startTime == "" || endTime == "" {
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "all fields are required",
			"Slot":  slot,
		})
		return
	}

	day, err := time.Parse("2006-01-02", dayStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "edit_slot.html", gin.H{
			"error": "invalid date format",
			"Slot":  slot,
		})
		return
	}

	slot.Day = day
	slot.StartTime = startTime
	slot.EndTime = endTime

	if err := h.repo.Update(&slot); err != nil {
		c.HTML(http.StatusInternalServerError, "edit_slot.html", gin.H{
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
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "invalid slot id",
		})
		return
	}

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(http.StatusNotFound, "slots.html", gin.H{
			"error": "slot not found",
		})
		return
	}

	if !slot.IsAvailable {
		c.HTML(http.StatusBadRequest, "slots.html", gin.H{
			"error": "cannot delete booked slot",
		})
		return
	}

	if err := h.repo.Delete(&model.TimeSlot{}, "id = ?", uint(slotID)); err != nil {
		c.HTML(http.StatusInternalServerError, "slots.html", gin.H{
			"error": "failed to delete slot",
		})
		return
	}

	c.Redirect(
		http.StatusFound,
		"/admin/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots",
	)
}
