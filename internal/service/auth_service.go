package service

import (
	"errors"
	"time"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

func RegisterUser(u *model.User) error {
	_, err := repository.FindUserByEmail(u.Email)
	if err == nil {
		return errors.New("email already registered")
	}

	u.Password, _ = utils.HashPassword(u.Password)
	return repository.CreateUser(u)
}

func LoginUser(email, password string) (model.User, string, string, time.Time, time.Time, error) {
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return model.User{}, "", "", time.Now(), time.Now(), errors.New("invalid email")
	}

	if !utils.CheckPassword(user.Password, password) {
		return model.User{}, "", "", time.Now(), time.Now(), errors.New("invalid password")
	}

	access, accessExp, _ := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	refresh, refreshExp, _ := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)
	rt := model.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		ExpiresAt: refreshExp,
	}

	repository.SaveRefreshToken(&rt)

	return user, access, refresh, accessExp, refreshExp, nil
}

func RefreshTokens(oldRefresh string) (string, string, time.Time, time.Time, error) {
	rt, err := repository.GetRefreshToken(oldRefresh)
	if err != nil {
		return "", "", time.Now(), time.Now(), errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {
		return "", "", time.Now(), time.Now(), errors.New("refresh token expired")
	}

	user, _ := repository.FindUserByID(rt.UserID)

	access, accessExp, _ := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	newRefresh, refreshExp, _ := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)

	rt.Token = newRefresh
	rt.ExpiresAt = refreshExp

	repository.UpdateRefreshToken(&rt)
	return access, newRefresh, accessExp, refreshExp, nil
}

func LogoutUser(token string) {
	repository.DeleteRefreshToken(token)
}
