package api

import "github.com/yl10/teambition/model"

// TBTag 认证接口
type TBTag interface {
	// GetTag 查询TAG
	GetTag(tagID string) (*model.Tag, error)

	// GetTagsOf 查询Tags
	GetTagsOf(objID, objType string) (*[]model.Tag, error)

	// AddTag 新增Tag
	AddTag(newTag *model.NewTag) (*model.Tag, error)

	// DeleteTag 删除Tag
	DeleteTag(TagID string) error
}
