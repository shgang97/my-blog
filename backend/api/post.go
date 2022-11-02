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

//func (api *Api) SaveAndUpdate(w http.ResponseWriter, r *http.Request) {
//	// 获取用户id， 判断用户是否登录
//	token := r.Header.Get("Authorization")
//	_, claim, err := util.ParseToken(token)
//	if err != nil {
//		common.Fail(w, err)
//		log.Println(err)
//		return
//	}
//	userId := claim.Uid
//
//	// POST save
//	method := r.Method
//	switch method {
//	case http.MethodPost:
//		// add
//		params := common.GetRequestJsonParam(r)
//		categoryId, _ := strconv.Atoi(params["CategoryId"].(string))
//		content := params["Content"].(string)
//		markdown := params["Markdown"].(string)
//		slug := params["Slug"].(string)
//		title := params["Title"].(string)
//		postType := int(params["Type"].(float64))
//		post := &models.Post{
//			Title:      title,
//			Content:    content,
//			Markdown:   markdown,
//			CategoryId: categoryId,
//			UserId:     userId,
//			ViewCount:  0,
//			Type:       postType,
//			Slug:       slug,
//			CreateAt:   time.Now(),
//			UpdateAt:   time.Now(),
//		}
//		pid := service.SavePost(post)
//		post.Pid = pid
//		common.Success(w, post)
//	case http.MethodPut:
//		// update
//		params := common.GetRequestJsonParam(r)
//		pid := int(params["Pid"].(float64))
//		categoryId, _ := strconv.Atoi(params["CategoryId"].(string))
//		content := params["Content"].(string)
//		markdown := params["Markdown"].(string)
//		slug := params["Slug"].(string)
//		title := params["Title"].(string)
//		postType := int(params["Type"].(float64))
//		viewCount := int(params["ViewCount"].(float64))
//		post := &models.Post{
//			Pid:        pid,
//			Title:      title,
//			Content:    content,
//			Markdown:   markdown,
//			CategoryId: categoryId,
//			UserId:     userId,
//			ViewCount:  viewCount,
//			Type:       postType,
//			Slug:       slug,
//			CreateAt:   time.Now(),
//			UpdateAt:   time.Now(),
//		}
//		service.UpdatePost(post)
//		common.Success(w, post)
//	}
//}
//
//func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
//	path := r.URL.Path
//	pid, err := strconv.Atoi(strings.TrimPrefix(path, "/api/post/"))
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	post, err := service.GetPostDetail(pid)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	common.Success(w, post)
//}
//
//func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	keywords := r.Form.Get("keywords")
//	searchList, err := service.SearchPost(keywords)
//	if err != nil {
//		log.Println(err)
//		common.Fail(w, err)
//	}
//	common.Success(w, searchList)
//}

func (*Api) Posts(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Index", err)
	}
	page, pageSize := 1, 10
	if pageStr := r.Form.Get("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Println("pageStr can not convert to int:", err)
			return
		}
	}
	if pageSizeStr := r.Form.Get("pageSize"); pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			log.Println("pageSizeStr can not convert to int:", err)
			return
		}
	}
	pagePost, err := service.Page(page, pageSize)
	if err != nil {
		log.Println("views.index.Index:", err)
		return
	}
	common.Success(w, pagePost)
}

func (*Api) Detail(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/post/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println("convert type string to type int failed:", err)
		return
	}
	post, err := service.GetPostDetail(pId)
	if err != nil {
		common.Fail(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) Edit(w http.ResponseWriter, r *http.Request) {
	// 获取用户id， 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := util.ParseToken(token)
	if err != nil {
		common.Fail(w, err)
		log.Println(err)
		return
	}
	userId := claim.Uid
	params := common.GetRequestJsonParam(r)
	content := params["content"].(string)
	title := params["title"].(string)
	description := params["description"].(string)
	post := &models.Post{
		UserId:      userId,
		Title:       title,
		Description: description,
		Content:     content,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		Status:      "0",
	}
	pid := service.SavePost(post)
	post.Id = pid
	common.Success(w, post)
}
