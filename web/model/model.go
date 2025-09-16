package model

import (
	"database/sql/driver"
	"encoding/json"
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

// Value 写库：用“值接收者”
func (t *LocalTime) Value() (driver.Value, error) {
	tt := time.Time(*t)
	if tt.IsZero() {
		return nil, nil
	}
	return tt, nil // 也可返回格式化字符串
}

// Scan 读库：用“指针接收者”
func (t *LocalTime) Scan(src any) error {
	if src == nil {
		*t = LocalTime(time.Time{})
		return nil
	}
	switch v := src.(type) {
	case time.Time:
		*t = LocalTime(v)
		return nil
	case []byte:
		s := string(v)
		if tm, err := time.ParseInLocation(time.RFC3339, s, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		if tm, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		if tm, err := time.ParseInLocation("2006-01-02 15:04:05.000", s, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		return fmt.Errorf("unsupported []byte time format: %q", s)
	case string:
		if tm, err := time.ParseInLocation(time.RFC3339, v, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		if tm, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		if tm, err := time.ParseInLocation("2006-01-02 15:04:05.000", v, time.Local); err == nil {
			*t = LocalTime(tm)
			return nil
		}
		return fmt.Errorf("unsupported string time format: %q", v)
	default:
		return fmt.Errorf("cannot scan %T into LocalTime", src)
	}
}
