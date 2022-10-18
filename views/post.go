package views

import (
	"errors"
	"log"
	"my-blog/common"
	"my-blog/service"
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
	//pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println("convert type string to type int failed:", err)
		detail.WriteError(w, err)
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		log.Println(err)
		detail.WriteError(w, errors.New("system error due to category, please contact the administrator"))
		return
	}
	detail.WriteData(w, postRes)
}
