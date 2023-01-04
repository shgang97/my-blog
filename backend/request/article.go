package request

import "backend/model"

/*
@author: shg
@since: 2022/12/13
@desc: //TODO
*/

type ArticleRequest struct {
	Article  model.Article  `form:"article"`
	Tags     []model.Tag    `from:"tags"`
	Category model.Category `form:"category"`
}
