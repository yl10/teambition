package services

import (
	"encoding/json"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultTagService 默认 TBTag 接口实现
type DefaultTagService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultTagService 创建实例
func NewDefaultTagService(cli *util.HttpClient, apiURL string) (*DefaultTagService, error) {
	tag := &DefaultTagService{cli: cli, apiURL: apiURL}
	return tag, nil
}

// GetTag 获取Tag
func (t *DefaultTagService) GetTag(tagID string) (*model.Tag, error) {
	if tagID == "" {
		return nil, util.NewError("查询Tag失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/tags/" + tagID

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, err
	}

	var tag model.Tag
	if err = json.Unmarshal(res, &tag); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &tag, nil
}

// GetTagsOf 获取Tag
func (t *DefaultTagService) GetTagsOf(objID, objType string) (*[]model.Tag, error) {
	if objID == "" || objType == "" {
		return nil, util.NewError("查询Tags失败, 没有指定类型与ID")
	}

	reqURL := t.apiURL + "api/" + objType + "/" + objID + "/tags"
	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, err
	}

	var tags []model.Tag
	if err = json.Unmarshal(res, &tags); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &tags, nil
}

// AddTag 新增Tag
func (t *DefaultTagService) AddTag(newTag *model.NewTag) (*model.Tag, error) {
	reqURL := t.apiURL + "api/tags"
	res, err := t.cli.Post(reqURL, newTag)
	if err != nil {
		return nil, err
	}

	var tag model.Tag
	if err = json.Unmarshal(res, &tag); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &tag, nil
}

// DeleteTag 删除Tag
func (t *DefaultTagService) DeleteTag(tagID string) error {
	if tagID == "" {
		return util.NewError("删除Tag失败, 没有TagID")
	}

	reqURL := t.apiURL + "api/tags/" + tagID

	_, err := t.cli.Delete(reqURL, nil)
	return err
}
