package response

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
