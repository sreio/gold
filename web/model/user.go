package model

type User struct {
	ComID
	Name    string `gorm:"column:name;type:string;size:256;not null;comment:用户名称" json:"name"`
	Cron    string `gorm:"column:cron;type:string;size:256;not null;comment:cron定时任务" json:"cron"`
	SaveDay int    `gorm:"column:save_day;type:int;not null;comment:推送日志保留天数" json:"save_day"`
	ComTime
}

type ApiUserConf struct {
	ComID
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ApiAddUser struct {
	Name    string `json:"name"`
	Cron    string `json:"cron"`
	SaveDay int    `json:"save_day"`
}

type ApiRequestListUser struct {
	ApiAddUser
	UserConf []UserConf `json:"user_conf" gorm:"foreignKey:UserID"`
}

type ApiRequestAddUser struct {
	ApiAddUser
	UserConf []ApiUserConf `json:"user_conf"`
}
