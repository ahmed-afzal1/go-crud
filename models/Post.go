package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
