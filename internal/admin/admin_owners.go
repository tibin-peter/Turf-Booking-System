package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

// List all owners
func (h *AdminHandler) ListOwners(c *gin.Context) {
	var owners []model.User

	if err := h.repo.FindMany(&owners, "role = ?", "owner"); err != nil {
		c.HTML(500, "admin_owners.html", gin.H{
			"error": "Failed to load owners",
		})
		return
	}

	c.HTML(200, "admin_owners.html", gin.H{
		"Owners": owners,
	})
}

// Show add owner page
func (h *AdminHandler) ShowAddOwnerPage(c *gin.Context) {
	c.HTML(200, "admin_addowners.html", nil)
}

// Create owner
func (h *AdminHandler) CreateOwner(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if name == "" || email == "" || password == "" {
		c.HTML(400, "admin_addowners.html", gin.H{
			"error": "All fields are required",
		})
		return
	}

	hashed, _ := utils.HashPassword(password)

	owner := model.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     "owner",
	}

	if err := h.repo.Insert(&owner); err != nil {
		c.HTML(500, "admin_addowners.html", gin.H{
			"error": "Email already exists",
		})
		return
	}

	c.Redirect(302, "/admin/owners")
}

// Block owner
func (h *AdminHandler) BlockOwner(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("owner_id"))

	var owner model.User
	if err := h.repo.FindById(&owner, uint(id)); err == nil {
		owner.IsBlocked = true
		h.repo.Update(&owner)
	}

	c.Redirect(302, "/admin/owners")
}

// Unblock owner
func (h *AdminHandler) UnblockOwner(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("owner_id"))

	var owner model.User
	if err := h.repo.FindById(&owner, uint(id)); err == nil {
		owner.IsBlocked = false
		h.repo.Update(&owner)
	}

	c.Redirect(302, "/admin/owners")
}
