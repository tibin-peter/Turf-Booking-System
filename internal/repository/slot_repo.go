package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func GetSlotByTurfID(turfID uint) ([]model.TimeSlot, error) {
	var slots []model.TimeSlot
	err := config.DB.Where("turf_id = ? ", turfID).Find(&slots).Error
	return slots, err
}
