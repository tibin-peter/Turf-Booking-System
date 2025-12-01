package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func GetAllTurfs() ([]model.Turf, error) {
	var turfs []model.Turf
	err := config.DB.Find(&turfs).Error
	return turfs, err
}

func GetTurfByID(id uint) (model.Turf, error) {
	var turf model.Turf
	err := config.DB.First(&turf, id).Error
	return turf, err
}
