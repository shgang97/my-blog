package views

import (
	"my-blog/common"
	"my-blog/config"
	"my-blog/models"
	"net/http"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

func (h *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 页面上涉及的所有数据定义
	var category = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	categories := []models.Category{
		{
			Cid:      1,
			Name:     "hello",
			Slug:     "world",
			CreateAt: "2022-11-11",
			UpdateAt: "2022-11-11",
		},
	}
	var posts = []models.PostMore{
		{
			Id:                   "000",
			Title:                "hello world",
			Slug:                 "slug",
			Description:          "my first blog",
			Cover:                "5",
			Markdown:             "6",
			Content:              "7",
			CopyrightType:        0,
			Original:             "8",
			OriginalAvatar:       "9",
			OriginalTitle:        "10",
			OriginalLink:         "11",
			Categories:           categories,
			ViewCount:            0,
			Type:                 0,
			CreationTime:         "12",
			LastModificationTime: "13",
		},
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
	var homeResponse = &models.HomeResponse{
		Viewer:   viewer,
		Category: category,
		Posts:    posts,
		Total:    1,
		Page:     1,
		Pages:    []int{1},
		PageEnd:  true,
	}
	index.WriteData(w, homeResponse)
}
