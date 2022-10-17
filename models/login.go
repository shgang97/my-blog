package models

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

type LoginRes struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
