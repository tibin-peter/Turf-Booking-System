package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// Insert all booking by user
func CreateBooking(b *model.Booking) error {
	return config.DB.Create(b).Error
}

// List all bookings by user
func GetUserBookings(userID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

// Get booking byID
func GetBookingByID(id uint) (model.Booking, error) {
	var booking model.Booking
	err := config.DB.First(&booking, id).Error
	return booking, err
}

// Update booking
func UpdateBooking(b *model.Booking) error {
	return config.DB.Save(b).Error
}
