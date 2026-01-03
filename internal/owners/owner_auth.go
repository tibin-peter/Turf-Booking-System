package owners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type OwnerHandler struct {
	repo repository.Repository
}

func NewOwnerHandler(repo repository.Repository) *OwnerHandler {
	return &OwnerHandler{repo: repo}
}

// Show login page
func (h *OwnerHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "owner_login.html", nil)
}

// handle login
func (h *OwnerHandler) OwnerLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var owner model.User

	err := h.repo.FindOne(&owner, "email = ?", email)
	if err != nil || owner.Role != "owner" {
		c.HTML(http.StatusUnauthorized, "owner_login.html", gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(owner.Password, password) {
		c.HTML(http.StatusUnauthorized, "owner_login.html", gin.H{
			"error": "Wrong password",
		})
		return
	}
	token, _, _ := utils.GenerateAccessToken(owner.ID, owner.Email, owner.Role)

	c.SetCookie("access_token", token, 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, "/owner/dashboard")
}

// logout
func (h *OwnerHandler) OwnerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/owner/login")
}
