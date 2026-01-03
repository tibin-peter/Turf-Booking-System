package owners

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func (h *OwnerHandler) ListMyTurfs(c *gin.Context) {

	ownerID := c.GetUint("user_id")

	var turfs []model.Turf
	if err := h.repo.FindMany(&turfs, "owner_id = ?", ownerID); err != nil {
		c.HTML(500, "owner_turfs.html", gin.H{
			"error": "failed to load turfs",
		})
		return
	}

	c.HTML(200, "owner_turfs.html", gin.H{
		"Turfs": turfs,
	})
}
func (h *OwnerHandler) ListSlots(c *gin.Context) {

	turfID, _ := strconv.Atoi(c.Param("id"))
	ownerID := c.GetUint("user_id")

	// ownership check
	var turf model.Turf
	if err := h.repo.FindOne(
		&turf,
		"id = ? AND owner_id = ?",
		turfID, ownerID,
	); err != nil {
		c.HTML(403, "owner_slots.html", gin.H{
			"error": "unauthorized turf access",
		})
		return
	}

	var slots []model.TimeSlot
	_ = h.repo.FindMany(&slots, "turf_id = ?", turfID)

	c.HTML(200, "owner_slots.html", gin.H{
		"TurfID": turfID,
		"Slots":  slots,
	})
}
func (h *OwnerHandler) AddSlot(c *gin.Context) {

	turfID, _ := strconv.Atoi(c.Param("id"))
	ownerID := c.GetUint("user_id")

	// ownership chekck
	var turf model.Turf
	if err := h.repo.FindOne(
		&turf,
		"id = ? AND owner_id = ?",
		turfID, ownerID,
	); err != nil {
		c.HTML(403, "owner_slots.html", gin.H{
			"error": "unauthorized",
		})
		return
	}

	dayStr := c.PostForm("day")
	start := c.PostForm("start_time")
	end := c.PostForm("end_time")

	day, _ := time.Parse("2006-01-02", dayStr)

	today := time.Now().Truncate(24 * time.Hour)
	if day.Before(today) {
		c.HTML(400, "owner_slots.html", gin.H{
			"error":  "past dates not allowed",
			"TurfID": turfID,
		})
		return
	}

	// overlap check
	count, _ := h.repo.Count(
		&model.TimeSlot{},
		`turf_id = ? AND day = ? AND 
		(start_time < ? AND end_time > ?)`,
		turfID, day, end, start,
	)
	if count > 0 {
		c.HTML(400, "owner_slots.html", gin.H{
			"error":  "slot already exists",
			"TurfID": turfID,
		})
		return
	}

	slot := model.TimeSlot{
		TurfID:    uint(turfID),
		Day:       day,
		StartTime: start,
		EndTime:   end,
	}

	_ = h.repo.Insert(&slot)

	c.Redirect(302, "/owner/turfs/"+strconv.Itoa(turfID)+"/slots")
}

func (h *OwnerHandler) ShowEditSlot(c *gin.Context) {

	slotID, _ := strconv.Atoi(c.Param("id"))
	ownerID := c.GetUint("user_id")

	var slot model.TimeSlot
	if err := h.repo.FindById(&slot, uint(slotID)); err != nil {
		c.HTML(404, "owner_editslot.html", gin.H{"error": "slot not found"})
		return
	}

	var turf model.Turf
	if err := h.repo.FindOne(&turf, "id = ? AND owner_id = ?", slot.TurfID, ownerID); err != nil {
		c.HTML(403, "owner_editslot.html", gin.H{"error": "unauthorized"})
		return
	}

	if !slot.IsAvailable {
		c.HTML(400, "owner_editslot.html", gin.H{"error": "slot already booked"})
		return
	}

	c.HTML(200, "owner_editslot.html", gin.H{"Slot": slot})
}
func (h *OwnerHandler) EditSlot(c *gin.Context) {

	slotID, _ := strconv.Atoi(c.Param("id"))

	var slot model.TimeSlot
	_ = h.repo.FindById(&slot, uint(slotID))

	slot.Day, _ = time.Parse("2006-01-02", c.PostForm("day"))
	slot.StartTime = c.PostForm("start_time")
	slot.EndTime = c.PostForm("end_time")

	_ = h.repo.Update(&slot)

	c.Redirect(302, "/owner/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots")
}
func (h *OwnerHandler) DeleteSlot(c *gin.Context) {

	slotID, _ := strconv.Atoi(c.Param("id"))

	var slot model.TimeSlot
	_ = h.repo.FindById(&slot, uint(slotID))

	if !slot.IsAvailable {
		c.Redirect(302, "/owner/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots")
		return
	}

	_ = h.repo.Delete(&model.TimeSlot{}, "id = ?", slotID)

	c.Redirect(302, "/owner/turfs/"+strconv.Itoa(int(slot.TurfID))+"/slots")
}
