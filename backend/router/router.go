package router

import (
	"backend/api/login"
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

	userRouter := authRouter.Group("/user")
	{
		userRouter.GET("/get", user.UserGet)
		userRouter.POST("/post", user.UserPost)
	}

	router.Run(":8080")
}
