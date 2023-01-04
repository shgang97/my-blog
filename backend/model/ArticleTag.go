package model

/*
@author: shg
@since: 2022/12/1
@desc: //TODO
*/

type ArticleTag struct {
	ArticleId string `json:"article_id"`
	TagId     string `json:"tag_id"`
	BaseModel
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
