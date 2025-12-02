package service

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

func ListSlotsByTurfID(turfID uint) ([]model.TimeSlot, error) {
	return repository.GetSlotByTurfID(turfID)
}
