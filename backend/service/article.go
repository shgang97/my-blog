package service

import (
	"backend/dao"
	"backend/model"
	"backend/request"
	"backend/response"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func GetAllArticleResponse(page, pageSize int) (*response.PageResult, error) {
	articles, err := dao.GetArticlesByPage(page, pageSize)
	total := dao.GetArticlesTotal()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responses []*response.ArticleResponse
	var ids []string
	for _, article := range articles {
		ids = append(ids, article.Id)
	}
	articleTagMap, _ := dao.GetTagByArticleIds(ids)
	articleCategoryMap, _ := dao.GetCategoryByArticleIds(ids)
	for _, article := range articles {
		articleTag := articleTagMap[article.Id]
		var tagIds, tagNames []string
		if articleTag != nil {
			tagIds = articleTag.TagIds
			tagNames = articleTag.TagNames
		}
		articleCategory := articleCategoryMap[article.Id]
		var categoryId, categoryName string
		if articleCategory != nil {
			categoryId = articleCategory.CategoryId
			categoryName = articleCategory.CategoryName
		}
		response := &response.ArticleResponse{
			Id:           article.Id,
			Title:        article.Title,
			Description:  article.Content,
			Cover:        article.Cover,
			ViewCount:    article.ViewCount,
			LikeCount:    article.LikeCount,
			CommentCount: article.CommentCount,
			TagIds:       tagIds,
			TagNames:     tagNames,
			CategoryId:   categoryId,
			CategoryName: categoryName,
		}
		responses = append(responses, response)
	}
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
