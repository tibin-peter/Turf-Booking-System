package service

import (
	"errors"

	"time"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type AuthService struct {
	repo repository.Repository
}

// dependency injection
func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(u *model.User) error {
	// Check if email exists
	var existing model.User
	err := s.repo.FindOne(&existing, "email = ?", u.Email)
	if err == nil {
		return errors.New("email already registered")
	}

	// Hash password
	u.Password, _ = utils.HashPassword(u.Password)

	// Insert user
	return s.repo.Insert(u)
}

func (s *AuthService) LoginUser(email, password string) (model.User, string, string, time.Time, time.Time, error) {
	var user model.User

	//find user by email
	err := s.repo.FindOne(&user, "email = ?", email)
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

	s.repo.Insert(&rt)

	return user, access, refresh, accessExp, refreshExp, nil
}

func (s *AuthService) RefreshTokens(oldRefresh string) (string, string, time.Time, time.Time, error) {
	var rt model.RefreshToken
	err := s.repo.FindOne(&rt, "token = ?", oldRefresh)
	if err != nil {
		return "", "", time.Now(), time.Now(), errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {
		return "", "", time.Now(), time.Now(), errors.New("refresh token expired")
	}

	var user model.User
	s.repo.FindById(&user, rt.UserID)

	access, accessExp, _ := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	newRefresh, refreshExp, _ := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)

	rt.Token = newRefresh
	rt.ExpiresAt = refreshExp

	s.repo.Update(&rt)
	return access, newRefresh, accessExp, refreshExp, nil
}

func (s *AuthService) LogoutUser(token string) {
	s.repo.Delete("token = ?", token)
}
