package api

import (
	"backend/common"
	"github.com/gin-gonic/gin"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func Logout(ctx *gin.Context) {
	common.Success(ctx, "logout success!")
}
