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
	ViewCount    string `json:"view_count"`
	LikeCount    string `json:"like_count"`
	CommentCount string `json:"comment_count"`
	BaseModel
}

func (Article) TableName() string {
	return "blog_article"
}
