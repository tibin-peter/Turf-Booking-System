package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	PassWord  string    `json:"_"`
	Role      string    `json:"role" gorm:"default:user"`
	CreatedAt time.Time `json:"created_at"`
}
