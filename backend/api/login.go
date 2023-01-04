package api

import (
	"backend/common"
	"backend/constant"
	request2 "backend/request"
	"backend/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

func Login(ctx *gin.Context) {
	var request request2.LoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		common.Fail(ctx, constant.ERROR_LOGIN_FAILED)
		return
	}

	response, err := service.Login(&request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Fail(ctx, constant.ERROR_USERNAME_OR_PASSWORD_NOT_EXIST)
		} else {
			common.Fail(ctx, constant.ERROR_LOGIN_FAILED)
		}
		return
	}
	common.Success(ctx, response)
}
