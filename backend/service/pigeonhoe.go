package service

import (
	"my-blog/backend/config"
	dao2 "my-blog/backend/dao"
	models2 "my-blog/backend/models"
)

/*
@author: shg
@since: 2022/10/21
@desc: //TODO
*/

func FindPostPigeonhole() (*models2.PigeonholeRes, error) {
	// 查询所有的文章 进行月份的整理
	posts, _ := dao2.GetAllPosts()
	pigeonholeMap := make(map[string][]models2.Post)
	for _, post := range posts {
		createTime := post.CreateAt
		month := createTime.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	// 查询所有的分类
	categries, _ := dao2.GetAllCategories()
	pigeonholeRes := &models2.PigeonholeRes{
		Viewer:     config.Config.Viewer,
		System:     config.Config.System,
		Categories: categries,
		Lines:      pigeonholeMap,
	}
	return pigeonholeRes, nil
}
