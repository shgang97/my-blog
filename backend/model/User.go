package model

/*
@author: shg
@since: 2022/11/30
@desc: //TODO
*/

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Ip       string `json:"ip"`
	BaseModel
}

// TableName 实现 Tabler 接口来更改默认表名
func (User) TableName() string {
	return "blog_user"
}
