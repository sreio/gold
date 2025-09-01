package core

import (
	"github.com/sreio/gold/config"
	"github.com/sreio/gold/database"
	"github.com/sreio/gold/logx"
)

func Init(cfg *config.Config) error {
	// 初始化日志
	logx.Init(&cfg.Logx)
	// 初始化数据库
	_, err := database.OpenDB(&cfg.DB)
	if err != nil {
		return err
	}
	// 初始化cron

	return nil
}
