package tag

import (
	"backend/common"
	"backend/constant"
	"backend/service/tag"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	tags, err := tag.GetAllTags()
	if err != nil {
		common.Fail(ctx, constant.ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE)
	}
	common.Success(ctx, tags)
}
