package model

import (
	"gorm.io/gorm"
	"time"
)

type ComID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type ComTime struct {
	CreatedAt time.Time      `gorm:"index;autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index;autoUpdateTime;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
