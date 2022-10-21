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
	// 根据文章id获取文章
	http.HandleFunc("/p/", views.HTML.Detail)
	// 写作页面
	http.HandleFunc("/writing/", views.HTML.Writing)
	// 归档页面
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	// 发布和更新文章
	http.HandleFunc("/api/post", api.API.SaveAndUpdate)
	// 获取发布后的文章
	http.HandleFunc("/api/post/", api.API.GetPost)
	//	上传图片
	http.HandleFunc("/api/qiniu/token", api.API.GetQiniuToken)
	//http.HandleFunc("/writing/?id", views.HTML.GetPost)
	http.HandleFunc("/api/login/account", api.API.Login)

	// 静态资源路由
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
