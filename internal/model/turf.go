package model

import "time"

type Turf struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	PricePerHour int       `json:"price_per_hour"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}
