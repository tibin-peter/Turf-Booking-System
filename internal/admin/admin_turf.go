package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// function for list all turfs
func (h *AdminHandler) AdminShowTurfs(c *gin.Context) {
	var turfs []model.Turf
	err := h.repo.FindMany(&turfs, "1=1")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "turfs_list.html", gin.H{
			"error": "Failed to load turfs",
		})
		return
	}

	c.HTML(http.StatusOK, "turfs_list.html", gin.H{
		"Turfs": turfs,
	})
}

// for showing the add page
func (h *AdminHandler) AdminShowAddTurfPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add_turf.html", nil)
}

// func for adding a new turf
func (h *AdminHandler) AdminAddTurf(c *gin.Context) {

	name := c.PostForm("name")
	location := c.PostForm("location")
	priceStr := c.PostForm("price")
	description := c.PostForm("description")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "add_turf.html", gin.H{
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
		c.HTML(http.StatusInternalServerError, "add_turf.html", gin.H{
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
		c.HTML(http.StatusNotFound, "turfs_list.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	c.HTML(http.StatusOK, "edit_turf.html", gin.H{"Turf": turf})
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
		c.HTML(http.StatusBadRequest, "edit_turf.html", gin.H{
			"error": "Invalid price",
		})
		return
	}

	if err := h.repo.FindById(&turf, uint(id)); err != nil {
		c.HTML(http.StatusNotFound, "edit_turf.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	turf.Name = name
	turf.Location = location
	turf.PricePerHour = price
	turf.Description = description

	if err := h.repo.Update(&turf); err != nil {
		c.HTML(http.StatusInternalServerError, "edit_turf.html", gin.H{
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

	h.repo.Delete("id = ?", uint(id))

	c.Redirect(http.StatusFound, "/admin/turfs")
}
