package router

import (
	"backend/api"
	"backend/api/category"
	"backend/api/tag"
	"backend/api/user"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

/*
@author: shg
@since: 2022/11/29
@desc: //TODO
*/

const ContextPath = "/api/blog"

func Router() {
	router := gin.New()
	router.Use(middleware.Cors())

	// TODO 不需要登录路由
	unauthRouter := router.Group(ContextPath)

	// TODO 需要登录路由
	authRouter := router.Group(ContextPath)
	authRouter.Use(middleware.JWT())

	unauthRouter.POST("/login", api.Login)
	authRouter.GET("logout", api.Logout)

	userRouter := authRouter.Group("/user")
	{
		userRouter.GET("/get", user.UserGet)
		userRouter.POST("/post", user.UserPost)
	}

	// 文章管理
	articleUnauthRouter := unauthRouter.Group("/articles")
	articleAuthRouter := authRouter.Group("/articles")
	{
		// 分页获取文章列表
		articleUnauthRouter.GET("", api.List)
		// 根据 id 阅读文章
		articleUnauthRouter.GET("/:id", api.Read)
		// 添加文章
		articleAuthRouter.POST("", api.Write)
		// 更新文章
		articleAuthRouter.PUT("/:id", api.Modify)
		// 删除文章
		articleAuthRouter.DELETE("/:id", api.Delete)
	}

	// 标签管理
	tagUnauthRouter := unauthRouter.Group("/tags")
	{
		tagUnauthRouter.GET("", tag.List)
	}

	// 标签管理
	categoryUnauthRouter := unauthRouter.Group("/categories")
	{
		categoryUnauthRouter.GET("", category.List)
	}

	_ = router.Run(":8080")
}
