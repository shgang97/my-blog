package category

import (
	"backend/dao"
	"backend/result"
)

func GetAllCategories() ([]*result.Category, error) {
	var categories []*result.Category
	dao.Db.Table("blog_category").
		Select("id, name").
		Find(&categories)
	return categories, nil
}
