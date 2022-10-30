package main

import (
	"fmt"
	_ "github.com/codyguo/godaemon"
	"log"
	"my-blog/backend/common"
	"my-blog/backend/router"
	"net/http"
)

/*
@author: shg
@since: 2022/10/12
@desc: //TODO
*/

func init() {
	// 模版加载
	common.LoadTemplate()
}
func main() {
	server := http.Server{
		Addr: ":8081",
	}
	// 路由
	router.Router()
	fmt.Println("blog system is running...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
