package repository

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"gorm.io/gorm"
)

type UserRepositoty struct {
	DB *gorm.DB
}

func NewUserRepositoty(db *gorm.DB) *UserRepositoty {
	return &UserRepositoty{DB: db}
}

func (r *UserRepositoty) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoty) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ? ", email).First(&user).Error
	return &user, err
}

func (r *UserRepositoty) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoty) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepositoty) Delete(id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}
