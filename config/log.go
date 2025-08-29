package config

type Logx struct {
	// 日志级别: "trace","debug","info","warn","error","fatal","panic"
	Level string `yaml:"level"`
	// 是否使用 JSON 格式（便于机器采集）；否则为人类友好的 text
	JSON bool `yaml:"json"`
	// 显示调用者（文件:行号），方便定位
	WithCaller bool `yaml:"withCaller"`
	// 控制台输出（stderr）
	Console bool `yaml:"console"`
	// 文件输出
	File struct {
		Enable     bool   `yaml:"enable"`
		Path       string `yaml:"path"`       // 如 logs/app.log
		MaxSize    int    `yaml:"maxSize"`    // 单个文件最大 MB
		MaxAge     int    `yaml:"maxAge"`     // 保留天数
		MaxBackups int    `yaml:"maxBackups"` // 备份文件个数
		Compress   bool   `yaml:"compress"`   // 是否 gzip 压缩旧文件
		// 如果只想把 >= Warn 的写入文件，可用 MinLevelForFile
		MinLevelForFile string `yaml:"minLevelForFile"` // 为空则与全局 Level 一致
	} `yaml:"file"`
	// 统一的时间格式（可选）
	TimeFormat string `yaml:"timeFormat"`
}
