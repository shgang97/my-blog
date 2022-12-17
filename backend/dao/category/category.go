package category

import (
	"backend/dao"
	"backend/model"
	"backend/result"
)

func GetAllCategories() ([]*result.Category, error) {
	var categories []*result.Category
	dao.Db.Table("blog_category").
		Select("id, name").
		Find(&categories)
	return categories, nil
}

func Insert(articleCategory *model.ArticleCategory) (string, error) {
	result := dao.Db.Create(articleCategory)
	if result.Error != nil {
		return "", result.Error
	}
	return articleCategory.Id, nil
}

func Update(articleCategory *model.ArticleCategory) {
	dao.Db.Table("blog_article_category").
		Select("category_id, update_at").
		Where("article_id = ?", articleCategory.ArticleId).
		Updates(articleCategory)
}
