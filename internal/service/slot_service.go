package service

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

// function for list slot by turf id
func ListSlotsByTurfID(turfID uint) ([]model.TimeSlot, error) {
	return repository.GetSlotsByTurfID(turfID)
}

// functon for list slot by date
func ListSlotByDate(turfID uint, date string) ([]model.TimeSlot, error) {
	return repository.GetSlotByTurfAndDate(turfID, date)
}
