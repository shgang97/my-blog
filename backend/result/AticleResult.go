package result

import "time"

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

type ArticleCategoryResult struct {
	Id           string
	Title        string
	Description  string
	Cover        string
	ViewCount    string
	LikeCount    string
	CommentCount string
	CategoryId   string
	CreateAt     time.Time
	UpdateAt     time.Time
}

type ArticleTag struct {
	ArticleId string
	TagIds    []string
	TagNames  []string
}

type ArticleCategory struct {
	ArticleId    string
	CategoryId   string
	CategoryName string
}
