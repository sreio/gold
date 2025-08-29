package model

type User struct {
	ComID
	Name    string `gorm:"column:name;type:string;size:256;not null;comment:用户名称" json:"name"`
	Cron    string `gorm:"column:cron;type:string;size:256;not null;comment:cron定时任务" json:"cron"`
	SaveDay int    `gorm:"column:save_day;type:int;not null;comment:推送日志保留天数" json:"save_day"`
	ComTime
}
