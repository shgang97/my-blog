package service

import (
	"html/template"
	"log"
	"my-blog/config"
	"my-blog/dao"
	"my-blog/models"
)

/*
@author: shg
@since: 2022/10/19
@desc: //TODO
*/

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(pid)
	var postRes *models.PostRes
	if err != nil {
		return nil, err
	}
	var postMore models.PostMore
	categoryName := dao.GetCategoryNameById(post.Pid)
	userName := dao.GetUserNameById(post.UserId)
	postMore = models.PostMore{
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
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}
	postRes = &models.PostRes{
		Viewer:  config.Config.Viewer,
		System:  config.Config.System,
		Article: postMore,
	}
	return postRes, nil
}

func Write() (*models.WritingRes, error) {
	var wr models.WritingRes
	wr.Title = config.Config.Viewer.Title
	wr.CdnURL = config.Config.System.CdnUrl
	categories, err := dao.GetAllCategories()
	if err != nil {
		log.Println("Write:", err)
		return nil, err
	}
	wr.Categories = categories
	return &wr, nil
}

func SavePost(post *models.Post) int {
	return dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(keywords string) ([]models.SearchRes, error) {
	searchList, err := dao.GetPostsByContent(keywords)
	if err != nil {
		return nil, err
	}
	return searchList, nil
}
