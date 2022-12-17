package dao

import (
	"backend/model"
	result2 "backend/result"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func GetArticleCategoryByPage(page, pageSize int) ([]*result2.ArticleCategory, error) {
	var results []*result2.ArticleCategory
	Db.Table("blog_article").
		Select("blog_article.id, blog_article.title, blog_article.content, blog_article.cover, blog_article.view_count, blog_article.like_count, blog_article.comment_count, blog_article.update_at, blog_article.create_at, blog_category.id as category_id, blog_category.name as category_name").
		Joins("left join blog_article_category on blog_article.id = blog_article_category.article_id").
		Joins("left join blog_category on blog_category.id=blog_article_category.category_id").
		Order("blog_article.create_at").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&results)
	return results, nil
}

func GetArticlesTotal() int {
	var total int64
	Db.Model(&model.Article{}).
		Count(&total)
	return int(total)
}

func GetArticleTagByArticleIds(ids []string) (map[string][]*result2.Tag, error) {
	var results []map[string]interface{}
	Db.Table("blog_article_tag").
		Select("blog_article_tag.article_id, blog_article_tag.tag_id as tag_id, blog_tag.name as tag_name").
		Joins("left join blog_tag on blog_article_tag.tag_id=blog_tag.id").
		Where("blog_article_tag.article_id IN ?", ids).
		Find(&results)

	articleTagMap := map[string][]*result2.Tag{}
	for _, result := range results {
		articleId := result["article_id"].(string)
		tagId := result["tag_id"].(string)
		tagName := result["tag_name"].(string)
		tag := &result2.Tag{
			Id:   tagId,
			Name: tagName,
		}
		if tags, ok := articleTagMap[articleId]; ok {
			articleTagMap[articleId] = append(tags, tag)
		} else {
			articleTagMap[articleId] = []*result2.Tag{tag}
		}
	}
	return articleTagMap, nil
}

func Insert(article model.Article) (string, error) {
	result := Db.Create(&article)
	if result.Error != nil {
		return "", result.Error
	}
	return article.Id, nil
}

func GetArticleById(id string) (*result2.Article, error) {
	var article result2.Article
	Db.Table("blog_article").
		Select("blog_article.id, blog_article.user_id, blog_article.title, blog_article.content, blog_article.cover, blog_article.view_count, blog_article.like_count, blog_article.comment_count, blog_article.create_at, blog_article.update_at").
		Where("id = ?", id).
		Take(&article)
	return &article, nil
}

func GetCategoryByArticleId(articleId string) (*result2.Category, error) {
	var category result2.Category
	Db.Table("blog_category").
		Select("blog_category.id, blog_category.name").
		Joins("left join blog_article_category on blog_category.id=blog_article_category.category_id").
		Where("blog_article_category.article_id = ?", articleId).
		Scan(&category)
	return &category, nil
}

func GetTagsByArticleId(articleId string) ([]*result2.Tag, error) {
	var tags []*result2.Tag
	Db.Table("blog_tag").
		Select("blog_tag.id, blog_tag.name").
		Joins("left join blog_article_tag on blog_tag.id=blog_article_tag.tag_id").
		Where("blog_article_tag.article_id = ?", articleId).
		Scan(&tags)
	return tags, nil
}
