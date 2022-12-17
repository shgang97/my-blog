package response

import (
	"backend/result"
)

/*
@author: shg
@since: 2022/12/10
@desc: // 用于封装返回的json数据
*/

type ArticleInfo struct {
	Article  *result.Article  `json:"article"`
	Tags     []*result.Tag    `json:"tags"`
	Category *result.Category `json:"category"`
}

type ArticleDetail struct {
	Article  *result.Article  `json:"article"`
	Tags     []*result.Tag    `json:"tags"`
	Category *result.Category `json:"category"`
}
