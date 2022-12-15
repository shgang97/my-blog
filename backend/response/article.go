package response

import (
	"backend/model"
	"backend/result"
)

/*
@author: shg
@since: 2022/12/10
@desc: // 用于封装返回的json数据
*/

type ArticleResponse struct {
	Id           string          `json:"id"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Cover        string          `json:"cover"`
	ViewCount    int             `json:"view_count"`
	LikeCount    int             `json:"like_count"`
	CommentCount int             `json:"comment_count"`
	Tags         []*model.Tag    `json:"tags"`
	Category     *model.Category `json:"category"`
}

type ArticleDetail struct {
	Article  *result.Article  `json:"article"`
	Tags     []*result.Tag    `json:"tags"`
	Category *result.Category `json:"category"`
}
