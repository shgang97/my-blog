package api

import (
	"log"
	"my-blog/backend/common"
	"my-blog/backend/service"
	"net/http"
)

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码 返回 对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	password := params["password"].(string)
	loginRes, err := service.Login(userName, password)
	if err != nil {
		log.Println("api Login:", err)
		common.Fail(w, err)
	}
	common.Success(w, loginRes)
}
