package api

import (
	"backend/common"
	"backend/constant"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func List(ctx *gin.Context) {
	page, pageSize := 1, 10
	var err error
	if pageStr := ctx.Query("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			common.FailWithStatus(ctx, http.StatusBadRequest, constant.ERROR_ARTICLE_FALSE_PAGE_PARAM)
			return
		}
	}
	if pageSizeStr := ctx.Query("page_size"); pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			common.FailWithStatus(ctx, http.StatusBadRequest, constant.ERROR_ARTICLE_PARSE_PAGE_SIZE_PARAM)
			return
		}
	}
	pageResult, err := service.GetAllArticleResponse(page, pageSize)
	if err != nil {
		common.FailWithStatus(ctx, http.StatusBadRequest, constant.ERROR_ARTICLE_PARSE_PAGE_SIZE_PARAM)
	}
	common.Success(ctx, pageResult)
}

func Read(ctx *gin.Context) {

}

func Write(ctx *gin.Context) {

}

func Modify(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
