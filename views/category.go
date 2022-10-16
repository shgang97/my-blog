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
@since: 2022/10/16
@desc: //TODO
*/

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	category := common.Template.Category
	path := r.URL.Path
	categoryId, err := strconv.Atoi(strings.Trim(path, "/c/"))
	if err != nil {
		log.Println("convert type string to type int failed:", err)
		category.WriteError(w, err)
		return
	}
	page, pageSize := 1, 10
	if pageStr := r.Form.Get("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Println("pageStr can not convert to int:", err)
			category.WriteError(w, err)
			return
		}
	}
	if pageSizeStr := r.Form.Get("pageSize"); pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			log.Println("pageSizeStr can not convert to int:", err)
			category.WriteError(w, err)
			return
		}
	}
	categoryResponse, err := service.GetPostsByCategoryId(categoryId, page, pageSize)
	if err != nil {
		log.Println(err)
		category.WriteError(w, errors.New("system error due to category, please contact the administrator"))
		return
	}
	category.WriteData(w, categoryResponse)

}
