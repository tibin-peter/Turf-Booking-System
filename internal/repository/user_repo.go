package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func CreateUser(u *model.User) error {
	return config.DB.Create(u).Error
}

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func FindUserByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return user, err
}
