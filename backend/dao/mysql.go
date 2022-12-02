package dao

import (
	"backend/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

/*
@author: shg
@since: 2022/11/30
@desc: //TODO
*/

var Db *gorm.DB
var err error

func init() {
	username := config.Config.DataSource.UserName
	password := config.Config.DataSource.Password
	ip := config.Config.DataSource.IP
	port := config.Config.DataSource.Port
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/blog?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, ip, port,
	)
	Db, err = gorm.Open(mysql.Open(dns))

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	sqlDB, _ := Db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
