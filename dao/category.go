package dao

import (
	"log"
	"my-blog/models"
)

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

func GetAllCategories() ([]models.Category, error) {
	rows, err := Db.Query("select * from category")
	if err != nil {
		log.Println("failed to query from table category:", err)
		return nil, err
	}
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.Slug, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("failed to convert to category:", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, err
}

func GetCategoryNameById(id int) string {
	row := Db.QueryRow("select name from category where cid = ?", id)
	var categoryName string
	row.Scan(&categoryName)
	return categoryName
}
