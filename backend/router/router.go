package router

import (
	"backend/api/article"
	"backend/api/login"
	"backend/api/logout"
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

	unauthRouter.POST("/login", login.Login)
	authRouter.GET("logout", logout.Login)

	userRouter := authRouter.Group("/user")
	{
		userRouter.GET("/get", user.UserGet)
		userRouter.POST("/post", user.UserPost)
	}

	// 文章管理
	articleUnauthRouter := unauthRouter.Group("/article")
	articleAuthRouter := authRouter.Group("/article")
	{
		// 分页获取文章列表
		articleUnauthRouter.GET("", article.List)
		// 根据 id 阅读文章
		articleUnauthRouter.GET("/:id", article.Read)
		// 添加文章
		articleAuthRouter.POST("", article.Write)
		// 更新文章
		articleAuthRouter.PUT("/:id", article.Modify)
		// 删除文章
		articleAuthRouter.DELETE("", article.Delete)
	}

	_ = router.Run(":8080")
}
