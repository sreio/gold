package config

type DB struct {
	Driver          string  `yaml:"driver"`          // 数据库类型
	MaxOpenConns    int     `yaml:"maxOpenConns"`    // 最大连接数
	MaxIdleConns    int     `yaml:"maxIdleConns"`    // 最大空闲连接数
	ConnMaxLifetime string  `yaml:"connMaxLifetime"` // 连接最大生命周期
	Sqlite3         Sqlite3 `yaml:"sqlite3"`         // sqlite3
	Mysql           Mysql   `yaml:"mysql"`           // mysql
}

type Sqlite3 struct {
	Path string `yaml:"path"` // 数据库文件路径
}

type Mysql struct {
	Host     string `yaml:"host"`     // 数据库主机
	Port     int    `yaml:"port"`     // 数据库端口
	User     string `yaml:"user"`     // 数据库用户
	Password string `yaml:"password"` // 数据库密码
	Database string `yaml:"database"` // 数据库名称
}
