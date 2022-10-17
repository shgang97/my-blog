package models

import "time"

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
