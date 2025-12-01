package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/config"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

func SaveRefreshToken(rt *model.RefreshToken) error {
	return config.DB.Create(rt).Error
}

func GetRefreshToken(token string) (model.RefreshToken, error) {
	var rt model.RefreshToken
	err := config.DB.Where("token = ?", token).First(&rt).Error
	return rt, err
}

func UpdateRefreshToken(rt *model.RefreshToken) error {
	return config.DB.Save(rt).Error
}

func DeleteRefreshToken(token string) {
	config.DB.Where("token = ?", token).Delete(&model.RefreshToken{})
}
