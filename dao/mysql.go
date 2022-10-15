package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/blog_system?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("failed to connect, err: ", err)
		panic(err)
	}
	// 最大空闲连接数，默认2个最大空闲连接
	Db.SetMaxIdleConns(5)
	//最大连接数，默认不限制最大连接数
	Db.SetMaxOpenConns(100)
	// 连接最大存活时间
	Db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	Db.SetConnMaxIdleTime(time.Minute * 1)
	err = Db.Ping()
	if err != nil {
		log.Println("failed to connect, err: ", err)
		_ = Db.Close()
		panic(err)
	}
}
