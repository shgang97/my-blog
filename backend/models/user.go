package models

import "time"

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

type User struct {
	Id         int       `json:"id"`
	UserName   string    `json:"userName"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	LastLogin  time.Time `json:"lastLogin"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
