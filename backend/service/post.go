package service

import (
	"html/template"
	"log"
	"my-blog/backend/config"
	dao2 "my-blog/backend/dao"
	models2 "my-blog/backend/models"
)

/*
@author: shg
@since: 2022/10/19
@desc: //TODO
*/

func GetPostDetail(pid int) (*models2.PostRes, error) {
	post, err := dao2.GetPostByPid(pid)
	var postRes *models2.PostRes
	if err != nil {
		return nil, err
	}
	var postMore models2.PostMore
	categoryName := dao2.GetCategoryNameById(post.CategoryId)
	userName := dao2.GetUserNameById(post.UserId)
	postMore = models2.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		Markdown:     post.Markdown,
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models2.DateDay(post.CreateAt),
		UpdateAt:     models2.DateDay(post.UpdateAt),
	}
	postRes = &models2.PostRes{
		Viewer:  config.Config.Viewer,
		System:  config.Config.System,
		Article: postMore,
	}
	return postRes, nil
}

func Write() (*models2.WritingRes, error) {
	var wr models2.WritingRes
	wr.Title = config.Config.Viewer.Title
	wr.CdnURL = config.Config.System.CdnUrl
	categories, err := dao2.GetAllCategories()
	if err != nil {
		log.Println("Write:", err)
		return nil, err
	}
	wr.Categories = categories
	return &wr, nil
}

func SavePost(post *models2.Post) int {
	return dao2.SavePost(post)
}

func UpdatePost(post *models2.Post) {
	dao2.UpdatePost(post)
}

func SearchPost(keywords string) ([]models2.SearchRes, error) {
	searchList, err := dao2.GetPostsByContent(keywords)
	if err != nil {
		return nil, err
	}
	return searchList, nil
}
