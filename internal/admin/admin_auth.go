package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type AdminHandler struct {
	repo repository.Repository
}

func NewAdminHandler(repo repository.Repository) *AdminHandler {
	return &AdminHandler{repo: repo}
}

// Show login page
func (h *AdminHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// handle login
func (h *AdminHandler) AdminLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var admin model.User

	err := h.repo.FindOne(&admin, "email = ?", email)
	if err != nil || admin.Role != "admin" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(admin.Password, password) {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Wrong password",
		})
		return
	}

	c.SetCookie("admin_session", email, 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

// logout
func (h *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("admin_session", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin/login")
}
