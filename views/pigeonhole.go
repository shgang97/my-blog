package views

import (
	"my-blog/common"
	"my-blog/service"
	"net/http"
)

/*
@author: shg
@since: 2022/10/21
@desc: //TODO
*/

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole
	pigeonholeRes, _ := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
