package views

import (
	"my-blog/common"
	"my-blog/config"
	"net/http"
)

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Config.Viewer)
}
