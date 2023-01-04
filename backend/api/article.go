package api

import (
	"backend/common"
	"backend/constant"
	request2 "backend/request"
	"backend/service"
	"fmt"
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
	params := ctx.Params
	id, _ := params.Get("id")
	article, err := service.GetArticleById(id)
	if err != nil {
		common.FailWithStatus(ctx, http.StatusOK, constant.ERROR_ARTICLE_PARSE_PAGE_SIZE_PARAM)
	}
	common.Success(ctx, article)
}

func Write(ctx *gin.Context) {
	var request request2.ArticleRequest
	r := ctx.Request
	fmt.Println(r)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		common.Fail(ctx, constant.ERROR_ARTICLE_WRITING)
		return
	}
	id, err := service.Write(request)
	if err != nil {
		common.FailWithStatus(ctx, http.StatusNotAcceptable, constant.ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE)
	}
	common.Success(ctx, id)
}

func Modify(ctx *gin.Context) {
	var request request2.ArticleRequest
	r := ctx.Request
	fmt.Println(r)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		common.Fail(ctx, constant.ERROR_ARTICLE_WRITING)
		return
	}
	id, err := service.Modify(request)
	if err != nil {
		common.FailWithStatus(ctx, http.StatusNotAcceptable, constant.ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE)
	}
	common.Success(ctx, id)
}

func Delete(ctx *gin.Context) {

}
