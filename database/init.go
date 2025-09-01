package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	mysqlcfg "github.com/go-sql-driver/mysql"
	"github.com/sreio/gold/config"
	"github.com/sreio/gold/web/model"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

var DB *gorm.DB
var SqlDB *sql.DB

var err error
var driver string

var DbNotDriverError = errors.New("暂未实现此驱动，其他驱动可按需扩展")
var SqlDBError = errors.New("无法获取SqlDb")
var DbPingError = errors.New("数据库连接失败")
var SqliteCreateDirError = errors.New("创建目录失败")
var ParseMaxLeftTimeError = errors.New("解析最大连接时间失败")
var MysqlMissConfigError = errors.New("mysql 配置不完整")
var DbMigrateError = errors.New("数据库迁移失败")

func OpenDB(c *config.DB) (*gorm.DB, error) {
	log.Println("初始化数据库...")

	driver = strings.ToLower(c.Driver)

	switch driver {
	case "sqlite3":
		if err = os.MkdirAll("data", 0o755); err != nil {
			return nil, SqliteCreateDirError
		}
		DB, err = gorm.Open(
			sqlite.Open(c.Sqlite3.Path),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
				SkipDefaultTransaction:                   true, // 禁用默认事务
			},
		)

	case "mysql":
		// 校验必要字段
		if c.Mysql.Host == "" || c.Mysql.User == "" || c.Mysql.Database == "" {
			return nil, MysqlMissConfigError
		}
		port := c.Mysql.Port
		if port == 0 {
			port = 3306
		}

		// 用 go-sql-driver 的 Config 安全构造 DSN（避免特殊字符问题）
		mc := mysqlcfg.NewConfig()
		mc.User = c.Mysql.User
		mc.Passwd = c.Mysql.Password
		mc.Net = "tcp"
		mc.Addr = fmt.Sprintf("%s:%d", c.Mysql.Host, port)
		mc.DBName = c.Mysql.Database
		mc.Params = map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True", // 正确解析 time.Time
			"loc":       "Local",
			//"timeout":      "3s",
			//"readTimeout":  "5s",
			//"writeTimeout": "5s",
		}
		// 如需 TLS，可在 mc.TLSConfig 配置后 mc.Params["tls"] = "<name>"

		dsn := mc.FormatDSN()

		DB, err = gorm.Open(
			gormmysql.Open(dsn),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
				SkipDefaultTransaction:                   true, // 禁用默认事务
			},
		)

	default:
		return nil, DbNotDriverError
	}

	if err != nil {
		return nil, err
	}

	SqlDB, err = DB.DB()
	if err != nil {
		return nil, SqlDBError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = SqlDB.PingContext(ctx); err != nil {
		return nil, DbPingError
	}

	SqlDB.SetMaxOpenConns(c.MaxOpenConns) // 最大连接数
	SqlDB.SetMaxIdleConns(c.MaxIdleConns) // 最大空闲连接数
	duration, err := time.ParseDuration(c.ConnMaxLifetime)
	if err != nil {
		return nil, ParseMaxLeftTimeError
	}
	SqlDB.SetConnMaxLifetime(duration) // 连接最大生命周期

	// 数据库迁移
	if err = AutoMigrate(); err != nil {
		return nil, err
	}

	return DB, nil
}

func AutoMigrate() error {
	log.Println("数据库迁移...")
	err = DB.AutoMigrate(
		&model.User{},
		&model.UserConf{},
		&model.Gold{},
		&model.TaskLog{},
	)
	if err != nil {
		return DbMigrateError
	}
	return nil
}
