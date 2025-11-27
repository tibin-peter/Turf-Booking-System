package model

import "time"

type TimeSlot struct {
	ID          uint      `gorm:"primaryKey"`
	TurfID      uint      `json:"turf_id"`
	Day         time.Time `json:"day"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
}
