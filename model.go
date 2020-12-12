package ono

import (
	"fmt"
	"github.com/flxxyz/ono/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func init() {
	var err error
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
	db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			SkipInitializeWithVersion: false,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   cfg.Prefix,
				SingularTable: true, // 不使用复数表名
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

// Orm gorm实例
func Orm() *gorm.DB {
	return db
}

// Prefix 获取表前缀
func Prefix(name string) string {
	return db.NamingStrategy.TableName(name)
}
