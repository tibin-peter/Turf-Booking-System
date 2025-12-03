package service

import (
	"errors"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
	"github.com/tibin-peter/Turf-Booking-System/internal/utils"
)

// func for get the user profile
func GetUserProfile(UserID uint) (model.User, error) {
	return repository.FindUserByID(UserID)
}

// func for update the user profile
func UpdateUserProfile(userID uint, data model.User) error {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	//update name
	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Email != "" && data.Email != user.Email {
		//update email
		exists, _ := repository.FindUserByEmail(data.Email)
		if exists.ID != 0 {
			return errors.New("email alredy in use")
		}
		user.Email = data.Email
	}
	//update password
	if data.Password != "" {
		hased, _ := utils.HashPassword(data.Password)
		user.Password = hased
	}
	return repository.UpdateUser(&user)
}

// func for getting the booking history
func GetBookingHistory(userID uint) ([]model.Booking, error) {
	return repository.GetUserBookings(userID)
}
