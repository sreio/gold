package model

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `gorm:"primaryKey"`
}

type Time struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
