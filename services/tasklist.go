package services

import (
	"encoding/json"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultTasklistService 默认 TBTasklist 接口实现
type DefaultTasklistService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultTasklistService 创建实例
func NewDefaultTasklistService(cli *util.HttpClient, apiURL string) (*DefaultTasklistService, error) {
	tasklist := &DefaultTasklistService{cli: cli, apiURL: apiURL}
	return tasklist, nil
}

// GetTasklist 获取任务分组
func (t *DefaultTasklistService) GetTasklist(tasklistID string) (*model.Tasklist, error) {
	if tasklistID == "" {
		return nil, util.NewError("查询任务分组失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/tasklists/" + tasklistID

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, err
	}

	var tasklist model.Tasklist
	if err = json.Unmarshal(res, &tasklist); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &tasklist, nil
}

// AddTasklist 新增分组
func (t *DefaultTasklistService) AddTasklist(newTasklist *model.NewTasklist) (*model.Tasklist, error) {
	reqURL := t.apiURL + "api/tasklists"
	res, err := t.cli.Post(reqURL, newTasklist)
	if err != nil {
		return nil, err
	}

	var tasklist model.Tasklist
	if err = json.Unmarshal(res, &tasklist); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &tasklist, nil
}

// DeleteTasklist 删除分组
func (t *DefaultTasklistService) DeleteTasklist(tasklistID string) error {
	if tasklistID == "" {
		return util.NewError("删除分组失败, 没有分组ID")
	}

	reqURL := t.apiURL + "api/tasklists/" + tasklistID

	_, err := t.cli.Delete(reqURL, nil)
	return err
}
