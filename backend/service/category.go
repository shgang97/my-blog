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
@since: 2022/10/16
@desc: //TODO
*/

func GetPostsByCategoryId(categoryId, page, pageSize int) (*models2.CategoryResponse, error) {
	categories, err := dao2.GetAllCategories()
	if err != nil {
		log.Println("GetAllIndexInfo failed:", err)
		return nil, err
	}
	// 分页获取文章
	posts, err := dao2.GetPostsPageByCategoryId(categoryId, page, pageSize)
	if err != nil {
		log.Println("GetPostsPage failed:", err)
		return nil, err
	}
	var postMores []models2.PostMore
	for _, post := range posts {
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		categoryName := dao2.GetCategoryNameById(categoryId)
		userName := dao2.GetUserNameById(post.UserId)
		postMores = append(postMores, models2.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(post.Content),
			CategoryId:   categoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models2.DateDay(post.CreateAt),
			UpdateAt:     models2.DateDay(post.UpdateAt),
		})
	}
	viewer := config.Viewer{
		Title:       config.Config.Viewer.Title,
		Description: config.Config.Viewer.Description,
		Logo:        config.Config.Viewer.Logo,
		Navigation:  config.Config.Viewer.Navigation,
		Bilibili:    config.Config.Viewer.Bilibili,
		Avatar:      config.Config.Viewer.Avatar,
		UserName:    config.Config.Viewer.UserName,
		UserDesc:    config.Config.Viewer.UserDesc,
	}
	total := dao2.GetPostTotalByCategoryId(categoryId)
	var pages []int
	pageNum := ((total - 1) / pageSize) + 1
	for i := 0; i < pageNum; i++ {
		pages = append(pages, i+1)
	}
	var homeResponse = &models2.HomeResponse{
		Viewer:     viewer,
		Categories: categories,
		Posts:      postMores,
		Total:      total,
		Page:       page,
		Pages:      pages,
		PageEnd:    page != pageNum,
	}
	categoryName := dao2.GetCategoryNameById(categoryId)
	categoryResponse := &models2.CategoryResponse{
		HomeResponse: homeResponse,
		CategoryName: categoryName,
	}
	return categoryResponse, nil
}
