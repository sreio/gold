package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Logx `yaml:"logx"`
	DB   `yaml:"db"`
	Web  `yaml:"web"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	log.Println("加载配置文件:", configPath)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}
