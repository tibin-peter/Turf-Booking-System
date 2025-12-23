package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

//fun for list all users
func (h *AdminHandler) ListUsers(c *gin.Context) {
	var users []model.User

	if err := h.repo.FindMany(&users, "role != ?", "admin"); err != nil {
		c.HTML(http.StatusInternalServerError, "users.html", gin.H{
			"error": "failed to load users",
		})
		return
	}
	c.HTML(http.StatusOK, "users.html", gin.H{
		"Users": users,
	})
}

//func for block user
func (h *AdminHandler) BlockUser(c *gin.Context) {
	idStr := c.Param("user_id")
	id, _ := strconv.Atoi(idStr)

	var user model.User
	if err := h.repo.FindById(&user, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/admin/users")
		return
	}
	user.IsBlocked = true
	h.repo.Update(&user)

	c.Redirect(http.StatusFound, "/admin/users")
}

//func for unbolck user
func (h *AdminHandler) UnblockUser(c *gin.Context) {
	idStr := c.Param("user_id")
	id, _ := strconv.Atoi(idStr)

	var user model.User
	if err := h.repo.FindById(&user, uint(id)); err != nil {
		c.Redirect(http.StatusFound, "/admin/users")
		return
	}

	user.IsBlocked = false
	h.repo.Update(&user)

	c.Redirect(http.StatusFound, "/admin/users")
}
