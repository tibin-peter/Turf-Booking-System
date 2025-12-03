package model

import "time"

type Booking struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	TurfID        uint      `json:"turf_id" gorm:"not null"`
	SlotID        uint      `json:"slot_id" gorm:"not null"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"payment_method" gorm:"default:'cash'"`
	PaymentStatus string    `json:"payment_status" gorm:"default:'pending'"`
	Status        string    `json:"status" gorm:"default:'pending'"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
