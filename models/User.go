package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
