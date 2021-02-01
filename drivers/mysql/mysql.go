package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	DefaultHost            = "127.0.0.1"
	DefaultPort       uint = 3306
	DefaultCharset         = "utf8mb4"
	DefaultStringSize uint = 256
)

type MysqlConfig struct {
	Username          string        `json:"username"`            // 用户名
	Password          string        `json:"password"`            // 密码
	DBName            string        `json:"dbname"`              // 数据库名
	Host              string        `json:"host"`                // 地址
	Port              uint          `json:"port"`                // 端口
	MaxOpenConn       int           `json:"max_open_conn"`       // 最大连接数
	MaxIdleConn       int           `json:"max_idle_conn"`       // 最大空闲连接数
	ConnMaxLifeTime   time.Duration `json:"conn_max_life_time"`  // 连接最大存活时间
	Charset           string        `json:"charset"`             // 字符集 utf8mb4
	DefaultStringSize uint          `json:"default_string_size"` // string 长度默认值
}

func NewMysqlWithConfig(config *MysqlConfig) (*gorm.DB, error) {
	if config.Host == "" {
		config.Host = DefaultHost
	}
	if config.Port == 0 {
		config.Port = DefaultPort
	}
	if config.Charset == "" {
		config.Charset = DefaultCharset
	}
	if config.DefaultStringSize == 0 {
		config.DefaultStringSize = DefaultStringSize
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%sparseTime=true&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName, config.Charset)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: config.DefaultStringSize,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.MaxOpenConn > 0 {
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	}
	if config.MaxIdleConn > 0 {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	}
	if config.ConnMaxLifeTime > 0 {
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifeTime)
	}
	return db, nil
}
