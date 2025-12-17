package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/service"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}
func (h *AuthHandler) Register(c *gin.Context) {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		utils.JSONError(c, 400, "invalid input")
		return
	}

	if err := h.service.RegisterUser(&u); err != nil {
		utils.JSONError(c, 400, err.Error())
		return
	}

	utils.JSONSuccess(c, "registration successful", nil)
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

	user, access, refresh, accessExp, refreshExp, err :=
		h.service.LoginUser(body.Email, body.Password)

	if err != nil {
		utils.JSONError(c, 401, err.Error())
		return
	}

	utils.SetCookie(c, "access_token", access, accessExp)
	utils.SetCookie(c, "refresh_token", refresh, refreshExp)

	user.Password = ""
	utils.JSONSuccess(c, "login successful", gin.H{"user": user})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		utils.JSONError(c, 401, "refresh token missing")
		return
	}

	access, newRefresh, accessExp, refreshExp, err :=
		h.service.RefreshTokens(rt)

	if err != nil {
		utils.JSONError(c, 401, err.Error())
		return
	}

	utils.SetCookie(c, "access_token", access, accessExp)
	utils.SetCookie(c, "refresh_token", newRefresh, refreshExp)

	utils.JSONSuccess(c, "token refreshed", nil)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		utils.JSONError(c, 400, "refresh token missing")
		return
	}

	h.service.LogoutUser(rt)

	utils.ClearCookie(c, "access_token")
	utils.ClearCookie(c, "refresh_token")

	utils.JSONSuccess(c, "logout successful", nil)
}
