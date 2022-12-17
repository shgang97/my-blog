package service

import (
	"backend/dao"
	"backend/model"
	"backend/request"
	"backend/response"
	"backend/result"
	"github.com/google/uuid"
	"strings"
	"time"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func GetAllArticleResponse(page, pageSize int) (*response.PageResult, error) {
	articleCategories, _ := dao.GetArticleCategoryByPage(page, pageSize)

	var responses []*response.ArticleInfo
	var ids []string
	for _, value := range articleCategories {
		ids = append(ids, value.Id)
	}
	articleTags, _ := dao.GetArticleTagByArticleIds(ids)
	for _, value := range articleCategories {
		responses = append(responses, &response.ArticleInfo{
			Article: &result.Article{
				Id:           value.Id,
				UserId:       value.UserId,
				Title:        value.Title,
				Content:      value.Content,
				Cover:        value.Cover,
				ViewCount:    value.ViewCount,
				LikeCount:    value.LikeCount,
				CommentCount: value.CommentCount,
				CreateAt:     value.CreateAt,
				UpdateAt:     value.UpdateAt,
			},
			Tags: articleTags[value.Id],
			Category: &result.Category{
				Id:   value.CategoryId,
				Name: value.CategoryName,
			},
		})
		ids = append(ids, value.Id)
	}
	total := dao.GetArticlesTotal()
	pageResult := &response.PageResult{
		Data:  responses,
		Total: total,
	}
	return pageResult, nil
}

func Write(request request.ArticleRequest) (string, error) {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	article := model.Article{
		UserId:       "",
		Title:        request.Title,
		Content:      request.Content,
		Cover:        "",
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
		BaseModel: model.BaseModel{
			Id:       id,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
	}
	dao.Insert(article)
	return article.Id, nil
}

func GetArticleById(id string) (*response.ArticleDetail, error) {
	article, _ := dao.GetArticleById(id)
	category, _ := dao.GetCategoryByArticleId(id)
	tags, _ := dao.GetTagsByArticleId(id)
	articleRes := &response.ArticleDetail{
		Article:  article,
		Tags:     tags,
		Category: category,
	}
	return articleRes, nil
}
