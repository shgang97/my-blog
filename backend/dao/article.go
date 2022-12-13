package dao

import (
	"backend/model"
	result2 "backend/result"
	"log"
	"strings"
)

/*
@author: shg
@since: 2022/12/10
@desc: //TODO
*/

func GetWithTagAndCategoryByPage(page, pageSize int) ([]result2.ArticleCategoryResult, error) {
	rows, err := Db.Table("blog_article").
		Select("blog_article.id, blog_article.title, blog_article.content, blog_article.cover, blog_article.view_count, blog_article.like_count, blog_article.comment_count, blog_article.update_at, blog_article.create_at, blog_article_category.id").
		Joins("left join blog_article_category on blog_article.id = blog_article_category.article_id").
		Order("blog_article.create_at").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Rows()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	var results []result2.ArticleCategoryResult
	for rows.Next() {
		var result result2.ArticleCategoryResult
		_ = rows.Scan(&result.Id, &result.Title, &result.Description, &result.Cover, &result.ViewCount, &result.LikeCount, &result.CommentCount, &result.UpdateAt, &result.CreateAt, &result.CategoryId)
		results = append(results, result)
	}
	return results, nil
}

func GetArticlesByPage(page, pageSize int) ([]*model.Article, error) {
	var articles []*model.Article
	result := Db.Model(&model.Article{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func GetArticlesTotal() int64 {
	var total int64
	Db.Model(&model.Article{}).
		Count(&total)
	return total
}

func GetTagByArticleIds(ids []string) (map[string]*result2.ArticleTag, error) {
	var results []map[string]interface{}
	Db.Table("blog_article_tag").
		Select("blog_article_tag.article_id, GROUP_CONCAT(DISTINCT blog_article_tag.tag_id) as tag_ids, GROUP_CONCAT(DISTINCT blog_tag.name) as tag_names").
		Joins("left join blog_tag on blog_article_tag.tag_id = blog_tag.id").
		Where("blog_article_tag.article_id IN ?", ids).
		Group("blog_article_tag.article_id").
		Find(&results)

	articleTagMap := map[string]*result2.ArticleTag{}
	for _, result := range results {
		articleId := result["article_id"].(string)
		tagIds := strings.Split(result["tag_ids"].(string), ",")
		tagNames := strings.Split(result["tag_names"].(string), ",")
		articleTagMap[articleId] = &result2.ArticleTag{
			ArticleId: articleId,
			TagIds:    tagIds,
			TagNames:  tagNames,
		}
	}
	return articleTagMap, nil
}

func GetCategoryByArticleIds(ids []string) (map[string]*result2.ArticleCategory, error) {
	var results []map[string]interface{}
	Db.Table("blog_article_category").
		Select("blog_article_category.article_id, blog_article_category.category_id, blog_category.name").
		Joins("left join blog_category on blog_article_category.category_id=blog_category.id").
		Where("blog_article_category.article_id IN ?", ids).
		Find(&results)
	articleCategoryMap := map[string]*result2.ArticleCategory{}
	for _, result := range results {
		articleId := result["article_id"].(string)
		categoryId := result["category_id"].(string)
		categoryName := result["category_name"].(string)
		articleCategoryMap[articleId] = &result2.ArticleCategory{
			ArticleId:    articleId,
			CategoryId:   categoryId,
			CategoryName: categoryName,
		}
	}
	return articleCategoryMap, nil
}

func Insert(article model.Article) (string, error) {
	result := Db.Create(&article)
	if result.Error != nil {
		return "", result.Error
	}
	return article.Id, nil
}
