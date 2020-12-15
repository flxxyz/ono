package model

import (
	"database/sql"
	"fmt"
	"github.com/flxxyz/ono/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	gormDB        *gorm.DB
	gormCfg       *gorm.Config
	gormDialector gorm.Dialector
)

func init() {
	cfg := conf.GetDBCfg()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)

	SetDialector(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}))

	SetConfig(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix,
			SingularTable: true, // 不使用复数表名
		},
	})
}

// Init orm初始化
func Init() (sqlDB *sql.DB, err error) {
	gormDB, err = gorm.Open(gormDialector, gormCfg)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err = gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

// SetDialector 设置orm驱动
func SetDialector(dialector gorm.Dialector) {
	gormDialector = dialector
}

// SetConfig 设置配置
func SetConfig(cfg *gorm.Config) {
	gormCfg = cfg
}

// Orm gorm实例
func Orm() *gorm.DB {
	return gormDB
}

// Prefix 获取表前缀
func Prefix(name string) string {
	return gormDB.NamingStrategy.TableName(name)
}
