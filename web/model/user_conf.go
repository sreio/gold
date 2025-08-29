package model

type UserConf struct {
	ComID
	Userid uint   `gorm:"column:user_id;type:int;not null;comment:用户id;index;" json:"user_id"`
	Type   string `gorm:"column:type;type:string;size:256;not null;comment:推送渠道类型;uniqueIndex:index_utkv" json:"type"`
	Key    string `gorm:"column:key;type:string;size:256;not null;comment:推送渠道-字段;uniqueIndex:index_utkv" json:"key"`
	Value  string `gorm:"column:value;type:string;size:256;not null;comment:推送渠道-字段值;uniqueIndex:index_utkv" json:"value"`
	ComTime
}
