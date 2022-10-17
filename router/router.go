package router

import (
	"my-blog/api"
	"my-blog/views"
	"net/http"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

// Router 路由/**
func Router() {
	// 1. 页面 views 2. 数据（json） 3.静态资源
	// 首页
	http.HandleFunc("/", views.HTML.Index)
	// 登陆页
	http.HandleFunc("/login", views.HTML.Login)
	// 根据分类查询
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/api/post", api.API.SaveAndUpdate)
	http.HandleFunc("/api/login/account", api.API.Login)

	// 静态资源路由
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
