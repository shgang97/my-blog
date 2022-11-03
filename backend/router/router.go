package router

import (
	"my-blog/backend/api"
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
	//http.HandleFunc("/", cors(views.HTML.Index))
	// 根据分类查询
	//http.HandleFunc("/c/", cors(views.HTML.Category))
	//// 根据文章id获取文章
	//http.HandleFunc("/p/", cors(views.HTML.Detail))
	//// 写作页面
	//http.HandleFunc("/writing/", cors(views.HTML.Writing))
	//// 归档页面
	//http.HandleFunc("/pigeonhole", cors(views.HTML.Pigeonhole))
	// 发布和更新文章
	//http.HandleFunc("/api/post", cors(api.API.SaveAndUpdate))
	//// 获取发布后的文章
	//http.HandleFunc("/api/post/", cors(api.API.GetPost))
	//// 根据关键字搜索文章
	//http.HandleFunc("/api/search", cors(api.API.SearchPost))
	//	上传图片
	//http.HandleFunc("/api/qiniu/token", cors(api.API.GetQiniuToken))
	//http.HandleFunc("/writing/?id", views.HTML.GetPost)
	http.HandleFunc("/api/login", cors(api.API.Login))
	http.HandleFunc("/api/logout", cors(api.API.Logout))

	// 文章列表
	http.HandleFunc("/api/posts", cors(api.API.Posts))
	http.HandleFunc("/api/post/", cors(api.API.Detail))
	http.HandleFunc("/api/post/edit", cors(api.API.Edit))

	// 静态资源路由
	//http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")                             //允许请求方法
		w.Header().Add("Content-Type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}
