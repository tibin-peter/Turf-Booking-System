package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

//////////////User Related////////////

// function for get the total users
func CountUsers() int {
	var count int64
	config.DB.Model(&model.User{}).Count(&count)
	return int(count)
}

/////////////Turf Related////////////

// function for count the turfs list
func CountTurfs() int {
	var count int64
	config.DB.Model(&model.Turf{}).Count((&count))
	return int(count)
}

/////////////Slot Related/////////////

// function for count the slots
func CountSlots() int {
	var count int64
	config.DB.Model(&model.TimeSlot{}).Count(&count)
	return int(count)
}

/////////////Booking Related/////////////

// function for count the bookings
func CountBookings() int {
	var count int64
	config.DB.Model(&model.Booking{}).Count(&count)
	return int(count)
}

// function for count bookings by date
func CountBookingsByDate(date string) int {
	var count int64
	config.DB.Model(&model.Booking{}).Where("DATE(created_at)=?", date).
		Count(&count)
	return int(count)
}
