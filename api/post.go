package api

import (
	"log"
	"my-blog/common"
	"my-blog/models"
	"my-blog/service"
	"my-blog/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

func (api *Api) SaveAndUpdate(w http.ResponseWriter, r *http.Request) {
	// 获取用户id， 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := util.ParseToken(token)
	if err != nil {
		common.Fail(w, err)
		log.Println(err)
		return
	}
	userId := claim.Uid

	// POST save
	method := r.Method
	switch method {
	case http.MethodPost:
		// add
		params := common.GetRequestJsonParam(r)
		categoryId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     userId,
			ViewCount:  0,
			Type:       postType,
			Slug:       slug,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// update
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pid, err := strconv.Atoi(strings.TrimPrefix(path, "/api/post/"))
	if err != nil {
		log.Println(err)
		return
	}
	post, err := service.GetPostDetail(pid)
	if err != nil {
		log.Println(err)
		return
	}
	common.Success(w, post)
}
