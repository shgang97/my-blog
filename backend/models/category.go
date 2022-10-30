package models

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/

// Category 从数据库获取
type Category struct {
	Cid      int
	Name     string
	Slug     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
