package model

type TaskLog struct {
	ComID
	Uid    uint   `gorm:"column:uid;type:int;not null;comment:用户id;index" json:"uid"`
	Status int    `gorm:"column:status;type:int;not null;comment:推送状态 1:成功,2失败" json:"status"`
	ErrMsg string `gorm:"column:err_msg;type:string;size:256;comment:错误信息" json:"err_msg"`
	ComTime
}
