package service

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

func ListTurfs() ([]model.Turf, error) {
	turfs, err := repository.GetAllTurfs()
	return turfs, err
}

func GetTurfByID(id uint) (model.Turf, error) {
	turf, err := repository.GetTurfByID(id)
	return turf, err
}
