package model

/*
@author: shg
@since: 2022/12/1
@desc: //TODO
*/

type Comment struct {
	UserId       string `json:"user_id"`
	TargetId     string `json:"target_id"`
	Type         int    `json:"type"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	BaseModel
}

func (Comment) TableName() string {
	return "blog_comment"
}
