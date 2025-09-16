package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type ComID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type ComTime struct {
	CreatedAt LocalTime `gorm:"type:datetime(0);index;autoCreateTime;" json:"created_at"`
	UpdatedAt LocalTime `gorm:"type:datetime(0);index;autoUpdateTime;" json:"updated_at"`
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *LocalTime) UnmarshalJSON(b []byte) error {
	// 允许空字符串
	if string(b) == `""` || string(b) == "null" {
		*t = LocalTime(time.Time{})
		return nil
	}
	// 去引号
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tt, err := time.ParseInLocation(time.DateTime, s, time.Local)
	if err != nil {
		return err
	}
	*t = LocalTime(tt)
	return nil
}
