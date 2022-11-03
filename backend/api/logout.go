package api

import (
	"my-blog/backend/common"
	"net/http"
)

/*
@author: shg
@since: 2022/11/3
@desc: //TODO
*/

func (*Api) Logout(w http.ResponseWriter, r *http.Request) {
	common.Success(w, "logout success！！！")
}
