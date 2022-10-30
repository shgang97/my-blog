package models

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}
