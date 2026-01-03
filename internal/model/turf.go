package model

import "time"

type Turf struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"name"`
	OwnerID      uint      `json:"owner_id" gorm:"not null;index"`
	Location     string    `json:"location"`
	PricePerHour int       `json:"price_per_hour"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Turf) TableName() string {
	return "turfs"
}
