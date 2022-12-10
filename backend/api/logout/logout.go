package logout

import (
	"backend/common"
	"github.com/gin-gonic/gin"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func Login(ctx *gin.Context) {
	common.Success(ctx, "logout success!")
}
