package service

import (
	"errors"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

type UserService struct {
	repo repository.Repository
}

// dependency injection
func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

// func for get the user profile
func (s *UserService) GetUserProfile(UserID uint) (model.User, error) {
	var user model.User
	err := s.repo.FindById(&user, UserID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// func for update the user profile
func (s *UserService) UpdateUserProfile(userID uint, data model.User) error {
	var user model.User
	if err := s.repo.FindById(&user, userID); err != nil {
		return errors.New("user not found")
	}
	//update name
	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Email != "" && data.Email != user.Email {
		//update email
		var existing model.User
		err := s.repo.FindOne(&existing, "email = ?", data.Email)
		if err == nil && existing.ID != 0 {
			return errors.New("email already existing")
		}
		user.Email = data.Email
	}
	//update password
	if data.Password != "" {
		hased, _ := utils.HashPassword(data.Password)
		user.Password = hased
	}
	return s.repo.Update(&user)
}

// func for getting the booking history
func (s *UserService) GetBookingHistory(userID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	err := s.repo.FindMany(&bookings, "user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
