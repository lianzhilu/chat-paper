package database

import (
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func init() {
	rcg := config.GetRuntimeConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		rcg.MySQLConfig.MySQLUsername,
		rcg.MySQLConfig.MySQLPassword,
		rcg.MySQLConfig.MySQLHost,
		rcg.MySQLConfig.MySQLPort,
		rcg.MySQLConfig.MySQLDatabase,
	)
	var err error

	// init database
	// for gorm.Config, see https://gorm.io/zh_CN/docs/gorm_config.html
	_db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return _db
}
