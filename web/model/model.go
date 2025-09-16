package model

import (
	"time"
)

type ComID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type ComTime struct {
	CreatedAt time.Time `gorm:"type:datetime(0);index;autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(0);index;autoUpdateTime;" json:"updated_at"`
}
