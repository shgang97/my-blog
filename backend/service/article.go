package service

import (
	"backend/dao"
	"backend/dao/category"
	"backend/dao/tag"
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
	articleId := strings.ReplaceAll(uuid.New().String(), "-", "")
	articleCategoryId := strings.ReplaceAll(uuid.New().String(), "-", "")
	now := time.Now()

	article := &request.Article
	article.Id = articleId
	article.CreateAt = now
	article.UpdateAt = now

	tags := request.Tags
	var articleTags []*model.ArticleTag
	for _, value := range tags {
		articleTagId := strings.ReplaceAll(uuid.New().String(), "-", "")
		articleTag := &model.ArticleTag{
			ArticleId: articleId,
			TagId:     value.Id,
			BaseModel: model.BaseModel{
				Id:       articleTagId,
				CreateAt: now,
				UpdateAt: now,
			},
		}
		articleTags = append(articleTags, articleTag)
	}

	articleCategory := &model.ArticleCategory{
		ArticleId:  articleId,
		CategoryId: request.Category.Id,
		BaseModel: model.BaseModel{
			Pk:       0,
			Id:       articleCategoryId,
			CreateAt: now,
			UpdateAt: now,
		},
	}

	dao.Insert(article)
	tag.BulkInsert(articleTags)
	category.Insert(articleCategory)
	return article.Id, nil
}

func Modify(request request.ArticleRequest) (string, error) {
	now := time.Now()
	article := &request.Article
	article.UpdateAt = now
	dao.Update(article)
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
