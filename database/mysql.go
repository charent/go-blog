package database

import (
	"fmt"
	"go-blog/config"
	"go-blog/utils/log"
	"gorm.io/driver/mysql"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func init() {
	//连接到mysql数据库
	var mysqlConfig = config.Server.Mysql
	var url string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DBname)

	var err error
	MysqlDB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Error.Printf("mysql open error, message: %v", err)
	}

	if MysqlDB.Error != nil {
		log.Error.Printf("database error %v", MysqlDB.Error)
	} else {
		log.Info.Printf("open mysql database %s success!", config.Server.Mysql.DBname)
	}
}
