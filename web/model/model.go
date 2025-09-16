package model

import (
	"fmt"
	"time"
)

type ComID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type ComTime struct {
	CreatedAt *LocalTime `gorm:"type:datetime(0);index;autoCreateTime;" json:"created_at"`
	UpdatedAt *LocalTime `gorm:"type:datetime(0);index;autoUpdateTime;" json:"updated_at"`
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}
