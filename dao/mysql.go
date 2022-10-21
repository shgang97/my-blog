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

type MyDb struct {
	*sql.DB
}

var Db MyDb

func init() {
	var err error
	dataSourceName := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/blog_system?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("failed to connect, err: ", err)
		panic(err)
	}
	// 最大空闲连接数，默认2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("failed to connect, err: ", err)
		_ = db.Close()
		panic(err)
	}
	Db = MyDb{db}
}
