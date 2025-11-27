package model

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	BookingID     uint      `json:"booking_id"`
	Method        string    `json:"method" gorm:"default:cash"`
	PaymentStatus string    `json:"payment_status" gorm:"default:success"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
