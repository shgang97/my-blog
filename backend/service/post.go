package service

import (
	"my-blog/backend/dao"
	"my-blog/backend/models"
)

/*
@author: shg
@since: 2022/10/19
@desc: //TODO
*/

func GetPostDetail(pid int) (*models.Post, error) {
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		return nil, err
	}
	return post, nil
}

//	func Write() (*models2.WritingRes, error) {
//		var wr models2.WritingRes
//		wr.Title = config.Config.Viewer.Title
//		wr.CdnURL = config.Config.System.CdnUrl
//		categories, err := dao2.GetAllCategories()
//		if err != nil {
//			log.Println("Write:", err)
//			return nil, err
//		}
//		wr.Categories = categories
//		return &wr, nil
//	}
func SavePost(post *models.Post) int {
	return dao.SavePost(post)
}

//
//func UpdatePost(post *models2.Post) {
//	dao2.UpdatePost(post)
//}
//
//func SearchPost(keywords string) ([]models2.SearchRes, error) {
//	searchList, err := dao2.GetPostsByContent(keywords)
//	if err != nil {
//		return nil, err
//	}
//	return searchList, nil
//}

func Page(page, pageSize int) (*models.PagePost, error) {
	posts, err := dao.GetPostsByPage(page, pageSize)
	if err != nil {
		return nil, err
	}
	total := dao.GetPostTotal()
	pagePost := &models.PagePost{
		Posts:    posts,
		Page:     page,
		Total:    total,
		PageSize: pageSize,
	}
	return pagePost, nil
}
