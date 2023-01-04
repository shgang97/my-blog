package request

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

type BasePageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
