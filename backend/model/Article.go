package model

/*
@author: shg
@since: 2022/12/1
@desc: //TODO
*/

type Article struct {
	UserId       string `json:"user_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Cover        string `json:"cover"`
	ViewCount    int    `json:"view_count"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	BaseModel
}

func (Article) TableName() string {
	return "blog_article"
}
