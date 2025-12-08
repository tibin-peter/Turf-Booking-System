package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

// function for list all turfs
func AdminShowTurfs(c *gin.Context) {
	turfs, err := repository.GetAllTurfs()
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
func AdminShowAddTurfPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add_turf.html", nil)
}

// func for adding a new turf
func AdminAddTurf(c *gin.Context) {

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

	err = repository.CreateTurf(&newTurf)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "add_turf.html", gin.H{
			"error": "Failed to add turf",
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/turfs")
}

// for showing the edit page
func AdminShowEditTurfPage(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	turf, err := repository.GetTurfByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "turfs_list.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	c.HTML(http.StatusOK, "edit_turf.html", gin.H{"Turf": turf})
}

// fuction for edit the turf details
func AdminEditTurf(c *gin.Context) {

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

	turf, err := repository.GetTurfByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "edit_turf.html", gin.H{
			"error": "Turf not found",
		})
		return
	}

	turf.Name = name
	turf.Location = location
	turf.PricePerHour = price
	turf.Description = description

	repository.UpdateTurf(&turf)

	c.Redirect(http.StatusFound, "/admin/turfs")
}

// function for delete a turf
func AdminDeleteTurf(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	repository.DeleteTurf(uint(id))

	c.Redirect(http.StatusFound, "/admin/turfs")
}
