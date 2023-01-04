package model

/*
@author: shg
@since: 2022/12/1
@desc: //TODO
*/

type ArticleCategory struct {
	ArticleId  string `json:"article_id"`
	CategoryId string `json:"category_id"`
	BaseModel
}

func (ArticleCategory) TableName() string {
	return "blog_article_category"
}
