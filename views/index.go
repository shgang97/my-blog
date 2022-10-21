package views

import (
	"errors"
	"fmt"
	"log"
	"my-blog/common"
	"my-blog/service"
	"net/http"
	"strconv"
	"strings"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

func (h *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 页面上涉及的所有数据定义
	err := r.ParseForm()
	if err != nil {
		log.Println("Index", err)
		index.WriteError(w, err)
	}
	page, pageSize := 1, 10
	if pageStr := r.Form.Get("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Println("pageStr can not convert to int:", err)
			index.WriteError(w, err)
			return
		}
	}
	if pageSizeStr := r.Form.Get("pageSize"); pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			log.Println("pageSizeStr can not convert to int:", err)
			index.WriteError(w, err)
			return
		}
	}
	path := r.URL.Path
	fmt.Println("path = ", path)
	slug := ""
	if strings.Contains(path, "/slug/") {
		slug = strings.TrimPrefix(path, "/slug/")
	}
	homeResponse, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("views.index.Index:", err)
		index.WriteError(w, errors.New("system error, please contact the administrator"))
		return
	}
	index.WriteData(w, homeResponse)
}
