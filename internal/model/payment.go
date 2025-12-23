package model

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey"`
	BookingID uint      `gorm:"not null"`
	Amount    int       `json:"amount"`
	Status    string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
