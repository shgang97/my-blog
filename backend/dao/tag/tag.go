package tag

import (
	"backend/dao"
	"backend/result"
)

func GetAllTags() ([]*result.Tag, error) {
	var tags []*result.Tag
	dao.Db.Table("blog_tag").
		Select("id, name").
		Find(&tags)
	return tags, nil
}
