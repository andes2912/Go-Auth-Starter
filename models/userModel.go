package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID	uint `gorm:"primaryKey"`
	Nama string
	Phone string
	Email string `gorm:"unique"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}