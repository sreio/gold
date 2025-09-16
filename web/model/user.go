package model

import (
	"github.com/sreio/gold/web/dto"
	"gorm.io/gorm"
)

type User struct {
	ComID
	Name    string `gorm:"column:name;type:string;size:256;not null;comment:用户名称" json:"name"`
	Cron    string `gorm:"column:cron;type:string;size:256;not null;comment:cron定时任务" json:"cron"`
	SaveDay int    `gorm:"column:save_day;type:int;not null;comment:推送日志保留天数" json:"save_day"`
	ComTime

	// 关键：声明外键和引用；加上级联（删除/更新）
	UserConf []UserConf `json:"user_conf" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string { return "user" }

type UserConf struct {
	ComID
	UserID uint   `gorm:"column:user_id;type:int;not null;comment:用户id;index;uniqueIndex:idx_utk" json:"user_id"`
	Type   string `gorm:"column:type;type:string;size:256;not null;comment:推送渠道类型;uniqueIndex:idx_utk" json:"type"`
	Key    string `gorm:"column:key;type:string;size:256;not null;comment:推送渠道-字段;uniqueIndex:idx_utk" json:"key"`
	Value  string `gorm:"column:value;type:string;size:256;not null;comment:推送渠道-字段值;" json:"value"`
	ComTime
}

func (UserConf) TableName() string { return "user_conf" }

func FilterUsers(q dto.QueryUser) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q.Name != "" {
			db = db.Where("name LIKE ?", "%"+q.Name+"%")
		}
		// 若需要按 UserConf.type 过滤
		if q.Type != "" {
			db = db.Joins("LEFT JOIN user_conf uc ON uc.user_id = user.id").
				Where("uc.type = ?", q.Type)
		}
		return db
	}
}
