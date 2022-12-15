package tag

import (
	"backend/dao/tag"
	"backend/result"
)

func GetAllTags() ([]*result.Tag, error) {
	tags, _ := tag.GetAllTags()
	return tags, nil
}
