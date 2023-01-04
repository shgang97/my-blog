package common

import (
	"backend/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

func Success(ctx *gin.Context, data interface{}) {
	var result Result
	result.Code = 200
	result.Message = "success"
	result.Data = data
	ctx.JSON(http.StatusOK, result)
}

func Fail(ctx *gin.Context, errCode int) {
	var result Result
	result.Code = errCode
	result.Message = constant.GetErrMsg(errCode)
	ctx.JSON(http.StatusUnauthorized, result)
}

func FailWithStatus(ctx *gin.Context, statusCode, errCode int) {
	var result Result
	result.Code = errCode
	result.Message = constant.GetErrMsg(errCode)
	ctx.JSON(statusCode, result)
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
