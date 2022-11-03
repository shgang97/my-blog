package service

import (
	models2 "my-blog/backend/models"
)

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

func GetAllIndexInfo(slug string, page, pageSize int) (*models2.HomeResponse, error) {
	// 获取分类
	//categories, err := dao2.GetAllCategories()
	//if err != nil {
	//	log.Println("GetAllIndexInfo failed:", err)
	//	return nil, err
	//}
	//// 分页获取文章
	//posts, err := dao2.GetPostsPage(slug, page, pageSize)
	//if err != nil {
	//	log.Println("GetPostsPage failed:", err)
	//	return nil, err
	//}
	//var postMores []models2.PostMore
	//for _, post := range posts {
	//	content := []rune(post.Content)
	//	if len(content) > 100 {
	//		content = content[:100]
	//	}
	//	categoryName := dao2.GetCategoryNameById(post.CategoryId)
	//	userName := dao2.GetUserNameById(post.UserId)
	//	postMores = append(postMores, models2.PostMore{
	//		Pid:          post.Pid,
	//		Title:        post.Title,
	//		Slug:         post.Slug,
	//		Content:      template.HTML(post.Content),
	//		CategoryId:   post.CategoryId,
	//		CategoryName: categoryName,
	//		UserId:       post.UserId,
	//		UserName:     userName,
	//		ViewCount:    post.ViewCount,
	//		Type:         post.Type,
	//		CreateAt:     models2.DateDay(post.CreateAt),
	//		UpdateAt:     models2.DateDay(post.UpdateAt),
	//	})
	//}
	//viewer := config.Viewer{
	//	Title:       config.Config.Viewer.Title,
	//	Description: config.Config.Viewer.Description,
	//	Logo:        config.Config.Viewer.Logo,
	//	Navigation:  config.Config.Viewer.Navigation,
	//	Bilibili:    config.Config.Viewer.Bilibili,
	//	Avatar:      config.Config.Viewer.Avatar,
	//	UserName:    config.Config.Viewer.UserName,
	//	UserDesc:    config.Config.Viewer.UserDesc,
	//}
	//total := dao2.GetPostTotal1(slug)
	//var pages []int
	//pageNum := ((total - 1) / pageSize) + 1
	//for i := 0; i < pageNum; i++ {
	//	pages = append(pages, i+1)
	//}
	//var homeResponse = &models2.HomeResponse{
	//	Viewer:     viewer,
	//	Categories: categories,
	//	Posts:      postMores,
	//	Total:      total,
	//	Page:       page,
	//	Pages:      pages,
	//	PageEnd:    page != pageNum,
	//}
	//return homeResponse, nil
	return nil, nil
}
