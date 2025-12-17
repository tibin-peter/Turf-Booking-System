package service

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

type SlotService struct {
	repo repository.Repository
}

func NewSlotService(repo repository.Repository) *SlotService {
	return &SlotService{repo: repo}
}

// function for list slot by turf id
func (s *SlotService) ListSlotsByTurfID(turfID uint) ([]model.TimeSlot, error) {
	var slots []model.TimeSlot
	err := s.repo.FindMany(&slots, "turf_id = ?", turfID)
	if err != nil {
		return []model.TimeSlot{}, err
	}
	return slots, nil
}

// functon for list slot by date
func (s *SlotService) ListSlotByDate(turfID uint, date string) ([]model.TimeSlot, error) {
	var slots []model.TimeSlot
	err := s.repo.FindMany(&slots, "turf_id = ? AND date = ?", turfID, date)
	if err != nil {
		return []model.TimeSlot{}, err
	}
	return slots, nil
}
