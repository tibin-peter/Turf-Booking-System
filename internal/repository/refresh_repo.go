package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func SaveRefreshToken(r *model.RefreshToken) error {
	return config.DB.Create(r).Error
}

func GetRefreshToken(token string) (*model.RefreshToken, error) {
	var r model.RefreshToken
	err := config.DB.Where("token = ?", token).First(&r).Error
	return &r, err
}

func DeleteRefreshToken(token string) error {
	return config.DB.Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}
