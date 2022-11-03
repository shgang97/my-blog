package models

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
