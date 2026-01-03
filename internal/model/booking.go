package model

import "time"

type Booking struct {
	ID            uint       `gorm:"primaryKey"`
	UserID        uint       `gorm:"not null;index"`
	TurfID        uint       `gorm:"not null;index"`
	SlotID        uint       `gorm:"not null;index"`
	TotalAmount   int        `gorm:"default:0"`
	PaymentMethod string     `gorm:"type:varchar(20);default:'cash'"`
	PaymentStatus string     `gorm:"type:varchar(20);default:'pending'"`
	Status        string     `gorm:"type:varchar(20);default:'pending'"`
	CancelledBy   string     `gorm:"type:varchar(20)"`
	CancelledAt   *time.Time `json:"canceled_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
