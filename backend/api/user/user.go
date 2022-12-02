package user

import (
	"backend/dao"
	"backend/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@author: shg
@since: 2022/12/2
@desc: //TODO
*/

func UserGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get:welcome to my blog",
	})
}

func UserPost(ctx *gin.Context) {
	var user model.User
	dao.Db.First(&user)
	userJson, _ := json.Marshal(user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "post:welcome to my blog",
		"user": string(userJson),
	})
}
