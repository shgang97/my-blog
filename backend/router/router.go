package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@author: shg
@since: 2022/11/29
@desc: //TODO
*/

func Router() {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "welcome to my blog",
		})
	})

	router.Run(":8080")
}
