package model

import "time"

type TimeSlot struct {
	ID          uint      `gorm:"primaryKey"`
	TurfID      uint      `json:"turf_id"`
	Day         time.Time `json:"day"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
