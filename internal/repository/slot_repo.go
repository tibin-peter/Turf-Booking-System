package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// Get all time slot by the turfid
func GetSlotsByTurfID(turfID uint) ([]model.TimeSlot, error) {
	var slots []model.TimeSlot
	err := config.DB.Where("turf_id = ? ", turfID).Find(&slots).Error
	return slots, err
}

// Get a particular time slot by slotid
func GetSlotByID(id uint) (model.TimeSlot, error) {
	var slot model.TimeSlot
	err := config.DB.First(&slot, id).Error
	return slot, err
}

// Update the slot for the avilability changes
func UpdateSlot(slot *model.TimeSlot) error {
	return config.DB.Save(slot).Error
}
