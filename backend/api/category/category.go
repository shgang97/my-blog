package category

import (
	"backend/common"
	"backend/constant"
	"backend/service/category"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	categories, err := category.GetAllCategories()
	if err != nil {
		common.Fail(ctx, constant.ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE)
	}
	common.Success(ctx, categories)
}
