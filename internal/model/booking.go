package model

import "time"

type Booking struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	TurfID        uint      `gorm:"not null"`
	SlotID        uint      `gorm:"not null"`
	TotalAmount   int       `gorm:"default:0"`
	PaymentMethod string    `gorm:"type:varchar(20);default:'cash'"`
	PaymentStatus string    `gorm:"type:varchar(20);default:'pending'"`
	Status        string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
