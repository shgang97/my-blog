package response

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

type ArticleResponse struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Cover        string   `json:"cover"`
	ViewCount    string   `json:"view_count"`
	LikeCount    string   `json:"like_count"`
	CommentCount string   `json:"comment_count"`
	TagIds       []string `json:"tag_ids"`
	TagNames     []string `json:"tag_names"`
	CategoryId   string   `json:"category_id"`
	CategoryName string   `json:"category_name"`
}
