package dto

type UserConfDTO struct {
	Type  string `json:"type"  binding:"required"`
	Key   string `json:"key"   binding:"required"`
	Value string `json:"value" binding:"required"`
}

type CreateUserDTO struct {
	Name     string        `json:"name"     binding:"required"`
	Cron     string        `json:"cron"     binding:"required"`
	SaveDay  int           `json:"save_day" binding:"required,gte=0"`
	UserConf []UserConfDTO `json:"user_conf"`
}

type UpdateUserDTO struct {
	Name     *string       `json:"name"` // 可选：PATCH 场景
	Cron     *string       `json:"cron"`
	SaveDay  *int          `json:"save_day"`
	UserConf []UserConfDTO `json:"user_conf"` // 约定：整量同步（覆盖）
}

type QueryUser struct {
	Page int    `form:"page,default=1"`
	Size int    `form:"size,default=20"`
	Name string `form:"name"` // 模糊搜索
	Type string `form:"type"` // 过滤 UserConf.type
}
