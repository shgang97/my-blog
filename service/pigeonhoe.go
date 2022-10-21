package service

import (
	"my-blog/config"
	"my-blog/dao"
	"my-blog/models"
)

/*
@author: shg
@since: 2022/10/21
@desc: //TODO
*/

func FindPostPigeonhole() (*models.PigeonholeRes, error) {
	// 查询所有的文章 进行月份的整理
	posts, _ := dao.GetAllPosts()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		createTime := post.CreateAt
		month := createTime.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	// 查询所有的分类
	categries, _ := dao.GetAllCategories()
	pigeonholeRes := &models.PigeonholeRes{
		Viewer:     config.Config.Viewer,
		System:     config.Config.System,
		Categories: categries,
		Lines:      pigeonholeMap,
	}
	return pigeonholeRes, nil
}
