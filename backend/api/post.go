package api

import (
	"log"
	"my-blog/backend/common"
	"my-blog/backend/models"
	"my-blog/backend/service"
	"my-blog/backend/util"
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
		categoryId, _ := strconv.Atoi(params["CategoryId"].(string))
		content := params["Content"].(string)
		markdown := params["Markdown"].(string)
		slug := params["Slug"].(string)
		title := params["Title"].(string)
		postType := int(params["Type"].(float64))
		post := &models.Post{
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
		pid := service.SavePost(post)
		post.Pid = pid
		common.Success(w, post)
	case http.MethodPut:
		// update
		params := common.GetRequestJsonParam(r)
		pid := int(params["Pid"].(float64))
		categoryId, _ := strconv.Atoi(params["CategoryId"].(string))
		content := params["Content"].(string)
		markdown := params["Markdown"].(string)
		slug := params["Slug"].(string)
		title := params["Title"].(string)
		postType := int(params["Type"].(float64))
		viewCount := int(params["ViewCount"].(float64))
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     userId,
			ViewCount:  viewCount,
			Type:       postType,
			Slug:       slug,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
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

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	keywords := r.Form.Get("keywords")
	searchList, err := service.SearchPost(keywords)
	if err != nil {
		log.Println(err)
		common.Fail(w, err)
	}
	common.Success(w, searchList)
}
