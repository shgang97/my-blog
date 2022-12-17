package result

import "time"

/*
@author: shg
@since: 2022/12/10
@desc: // 用于封装从数据库查询的结果
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
	ArticleId string `json:"article_id"`
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
}

type ArticleCategory struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Cover        string    `json:"cover"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
	CategoryId   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

type Article struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Cover        string    `json:"cover"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
