package services

import (
	"encoding/json"

	"net/url"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultTaskService 默认 TBTask 接口实现
type DefaultTaskService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultTaskService 创建实例
func NewDefaultTaskService(cli *util.HttpClient, apiURL string) (*DefaultTaskService, error) {
	task := &DefaultTaskService{cli: cli, apiURL: apiURL}
	return task, nil
}

// GetTask 获取任务
func (t *DefaultTaskService) GetTask(taskID string, detailType ...string) (*model.Task, error) {
	if taskID == "" {
		return nil, util.NewError("查询任务失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/tasks/" + taskID

	params := make(url.Values)
	if len(detailType) > 0 && detailType[0] != "" {
		params.Add("detailType", detailType[0])
	}

	res, err := t.cli.Get(reqURL, params)
	if err != nil {
		return nil, err
	}

	var task model.Task
	if err = json.Unmarshal(res, &task); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &task, nil
}

// AddTask 新增任务
func (t *DefaultTaskService) AddTask(newTask *model.NewTask) (*model.Task, error) {
	reqURL := t.apiURL + "api/tasks"
	res, err := t.cli.Post(reqURL, newTask)
	if err != nil {
		return nil, err
	}

	var task model.Task
	if err = json.Unmarshal(res, &task); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &task, nil
}

// DeleteTask 删除任务
func (t *DefaultTaskService) DeleteTask(taskID string) error {
	if taskID == "" {
		return util.NewError("删除任务失败, 没有任务ID")
	}

	reqURL := t.apiURL + "api/tasks/" + taskID

	_, err := t.cli.Delete(reqURL, nil)
	return err
}
