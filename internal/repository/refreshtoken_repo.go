package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{DB: db}
}

func (r *RefreshTokenRepository) Store(token *model.RefreshToken) error {
	return r.DB.Create(token).Error
}

func (r *RefreshTokenRepository) Find(token string) (*model.RefreshToken, error) {
	var rt model.RefreshToken
	err := r.DB.Where("token = ?", token).Find(&rt).Error
	return &rt, err
}
func (r *RefreshTokenRepository) Delete(token string) error {
	return r.DB.Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}
