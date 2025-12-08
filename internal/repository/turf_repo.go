package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

// function for  get the all turf
func GetAllTurfs() ([]model.Turf, error) {
	var turfs []model.Turf
	err := config.DB.Find(&turfs).Error
	return turfs, err
}

// function for get a particular turf by id
func GetTurfByID(id uint) (model.Turf, error) {
	var turf model.Turf
	err := config.DB.First(&turf, id).Error
	return turf, err
}

// function for create new turf
func CreateTurf(t *model.Turf) error {
	err := config.DB.Create(t).Error
	return err
}

// function for update turf
func UpdateTurf(t *model.Turf) error {
	return config.DB.Save(t).Error
}

// function for delete turf
func DeleteTurf(id uint) error {
	err := config.DB.Delete(&model.Turf{}, id).Error
	return err
}
