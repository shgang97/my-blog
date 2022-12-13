package request

/*
@author: shg
@since: 2022/12/13
@desc: //TODO
*/

type ArticleRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags`
	Category string   `json:"category"`
}
