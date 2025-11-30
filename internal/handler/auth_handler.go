package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		utils.JSONError(c, 400, "invalid input")
		return
	}

	if err := h.Service.Regiser(&u); err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}
	utils.JSONSuccess(c, "Signup successfull", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.JSONError(c, 400, "invalid input")
		return
	}
	fmt.Println("email", body.Email)
	user, access, refresh, accessExp, refreshExp, err := h.Service.Login(body.Email, body.Password)
	if err != nil {
		utils.JSONError(c, 401, err.Error())
		return
	}
	utils.SetCookie(c, "access_token", access, accessExp)
	utils.SetCookie(c, "refresh_token", refresh, refreshExp)

	user.Password = ""
	utils.JSONSuccess(c, "login successfull", gin.H{"user": user})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		utils.JSONError(c, 401, "refresh token missing")
		return
	}

	newAccess, newRefresh, accessExp, refreshExp, err := h.Service.Rotate(rt)
	if err != nil {
		utils.JSONError(c, 401, err.Error())
		return
	}

	utils.SetCookie(c, "access_token", newAccess, accessExp)
	utils.SetCookie(c, "refresh_token", newRefresh, refreshExp)

	utils.JSONSuccess(c, "tokens refreshed", nil)
}

// LOGOUT
func (h *AuthHandler) Logout(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		utils.JSONError(c, 400, "refresh token missing")
		return
	}

	h.Service.Logout(rt)

	utils.ClearCookie(c, "access_token")
	utils.ClearCookie(c, "refresh_token")

	utils.JSONSuccess(c, "logout successful", nil)
}
