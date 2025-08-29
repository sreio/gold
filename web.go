package main

import (
	"github.com/sreio/gold/config"
	"github.com/sreio/gold/core"
	"log"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("config_example.yaml")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 初始化依赖
	err = core.Init(*cfg)
	if err != nil {
		log.Fatalf("初始化依赖失败: %v", err)
	}

	// 启动服务
	err = core.Start(*cfg)
	if err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
