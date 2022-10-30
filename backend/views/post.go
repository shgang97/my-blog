package views

import (
	"fmt"
	"log"
	"my-blog/backend/common"
	"my-blog/backend/service"
	"net/http"
	"strconv"
	"strings"
)

/*
@author: shg
@since: 2022/10/19
@desc: //TODO
*/

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println("convert type string to type int failed:", err)
		detail.WriteError(w, err)
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		log.Println("(*HTMLApi) Detail:", err)
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		errMsg := fmt.Sprintf("not exist post: pid[%d]", pId)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	detail.WriteData(w, postRes)
}
