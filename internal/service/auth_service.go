package service

import (
	"errors"
	"time"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type AuthService struct{}

func (s *AuthService) Regiser(u *model.User) error {
	exist, _ := repository.FindUserByEmail(u.Email)
	if exist.ID != 0 {
		return errors.New("email already registered")
	}
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	u.Role = "user"

	return repository.CreateUser(u)
}
func (s *AuthService) Login(email, password string) (*model.User, string, string, time.Time, time.Time, error) {
	u, err := repository.FindUserByEmail(email)
	if err != nil || u.ID == 0 {
		return nil, "", "", time.Time{}, time.Time{}, errors.New("invalid data")
	}
	if !utils.CheckPassword(password, u.Password) {
		return nil, "", "", time.Time{}, time.Time{}, errors.New("wrong password")
	}
	access, accessExp, err := utils.GenerateAccessToken(u.ID, u.Email, u.Role)
	if err != nil {
		return nil, "", "", time.Time{}, time.Time{}, err
	}
	refresh, refreshExp, err := utils.GenerateAccessToken(u.ID, u.Email, u.Role)
	if err != nil {
		return nil, "", "", time.Time{}, time.Time{}, err
	}

	rt := &model.RefreshToken{
		UserID:    u.ID,
		Token:     refresh,
		ExpiresAt: refreshExp,
	}
	repository.SaveRefreshToken(rt)
	return u, access, refresh, accessExp, refreshExp, nil
}
func (s *AuthService) Rotate(oldToken string) (string, string, time.Time, time.Time, error) {
	rt, err := repository.GetRefreshToken(oldToken)
	if err != nil || rt.ID == 0 {
		return "", "", time.Time{}, time.Time{}, errors.New("invalid refresh token")
	}
	if time.Now().After(rt.ExpiresAt) {
		return "", "", time.Time{}, time.Time{}, errors.New("token expired please login")
	}
	user, err := repository.FindUserById(rt.UserID)
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}

	newAccess, accessExp, _ := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	newRefresh, refreshExp, _ := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)

	repository.DeleteRefreshToken(oldToken)

	repository.SaveRefreshToken(&model.RefreshToken{
		UserID:    user.ID,
		Token:     newRefresh,
		ExpiresAt: refreshExp,
	})
	return newAccess, newRefresh, accessExp, refreshExp, nil
}
func (s *AuthService) Logout(refresh string) error {
	return repository.DeleteRefreshToken(refresh)
}
