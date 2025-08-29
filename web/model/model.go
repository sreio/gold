package model

import (
	"gorm.io/gorm"
	"time"
)

type ComID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type ComTime struct {
	CreatedAt time.Time      `gorm:"index;autoCreateTime;"`
	UpdatedAt time.Time      `gorm:"index;autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
