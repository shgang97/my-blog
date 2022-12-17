package tag

import (
	"backend/dao"
	"backend/model"
	"backend/result"
)

func GetAllTags() ([]*result.Tag, error) {
	var tags []*result.Tag
	dao.Db.Table("blog_tag").
		Select("id, name").
		Find(&tags)
	return tags, nil
}

func Insert(articleTag *model.ArticleTag) (string, error) {
	result := dao.Db.Create(articleTag)
	if result.Error != nil {
		return "", result.Error
	}
	return articleTag.Id, nil
}

func BulkInsert(articleTags []*model.ArticleTag) (int64, error) {
	result := dao.Db.Create(articleTags)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
