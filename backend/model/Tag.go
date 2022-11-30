package model

/*
@author: shg
@since: 2022/12/1
@desc: //TODO
*/

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BaseModel
}

func (Tag) TableName() string {
	return "blog_tag"
}
