package model

import "time"

type Booking struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `json:"user_id"`
	TurfID        uint      `json:"turf_id"`
	Date          time.Time `json:"date"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"payment_method" gorm:"default:cash"`
	PaymentStatus string    `json:"payment_status" gorm:"default:sucess"`
	Status        string    `json:"status" gorm:"default:booked"`
	CreatedAt     time.Time `json:"created_at"`
}
