package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

// function for list all turfs
func (h *AdminHandler) AdminShowTurfs(c *gin.Context) {
	//get the query from the context
	page := 1
	if p := c.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	limit := 5

	total, _ := h.repo.Count(&model.Turf{}, "")
	pagination := utils.NewPagination(page, limit, total)

	var turfs []model.Turf
	err := h.repo.FindManyPaginated(&turfs, "1=1", pagination.Limit, pagination.Offset)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_turfs.html", gin.H{
			"error": "Failed to load turfs",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_turfs.html", gin.H{
		"Turfs":      turfs,
		"Pagination": pagination,
	})
}

// for showing the add page
func (h *AdminHandler) AdminShowAddTurfPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_addturf.html", nil)
}

// func for adding a new turf
func (h *AdminHandler) AdminAddTurf(c *gin.Context) {

	name := c.PostForm("name")
	location := c.PostForm("location")
	priceStr := c.PostForm("price")
	description := c.PostForm("description")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_addturf.html", gin.H{
			"error": "Invalid price",
		})
		return
	}

	newTurf := model.Turf{
		Name:         name,
		Location:     location,
		PricePerHour: price,
		Description:  description,
	}

	err = h.repo.Insert(&newTurf)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_addturf.html", gin.H{
			"error": "Failed to add turf",
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/turfs")
}

// for showing the edit page
func (h *AdminHandler) AdminShowEditTurfPage(c *gin.Context) {
	var turf model.Turf
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := h.repo.FindById(&turf, uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "admin_turfs.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_editturf.html", gin.H{"Turf": turf})
}

// fuction for edit the turf details
func (h *AdminHandler) AdminEditTurf(c *gin.Context) {
	var turf model.Turf
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	name := c.PostForm("name")
	location := c.PostForm("location")
	priceStr := c.PostForm("price")
	description := c.PostForm("description")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_editturf.html", gin.H{
			"error": "Invalid price",
		})
		return
	}

	if err := h.repo.FindById(&turf, uint(id)); err != nil {
		c.HTML(http.StatusNotFound, "admin_editturf.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	turf.Name = name
	turf.Location = location
	turf.PricePerHour = price
	turf.Description = description

	if err := h.repo.Update(&turf); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_editturf.html", gin.H{
			"error": "Failed to update turf",
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/turfs")
}

// function for delete a turf
func (h *AdminHandler) AdminDeleteTurf(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var turfs []model.Turf
	err := h.repo.FindMany(&turfs, "1 = 1")
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_turfs.html", gin.H{
			"error": "invalid turf id",
			"Turfs": turfs,
		})
		return
	}

	// Prevent deleting turf with slots
	var slots []model.TimeSlot
	_ = h.repo.FindMany(&slots, "turf_id = ?", uint(id))
	if len(slots) > 0 {
		c.HTML(http.StatusBadRequest, "admin_turfs.html", gin.H{
			"error": "cannot delete turf with existing slots",
			"Turfs": turfs,
		})
		return
	}

	if err := h.repo.Delete(&model.Turf{}, "id = ?", uint(id)); err != nil {
		c.HTML(http.StatusInternalServerError, "admin_turfs.html", gin.H{
			"error": "failed to delete turf",
			"Turfs": turfs,
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/turfs")
}
