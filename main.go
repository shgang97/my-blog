package main

import (
	"fmt"
	_ "github.com/codyguo/godaemon"
	"log"
	"my-blog/common"
	"my-blog/router"
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
		Addr: ":8080",
	}
	// 路由
	router.Router()
	fmt.Println("blog system is running...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}