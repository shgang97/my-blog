package category

import (
	"backend/dao/category"
	"backend/result"
)

func GetAllCategories() ([]*result.Category, error) {
	categories, _ := category.GetAllCategories()
	return categories, nil
}
