package api

import "github.com/yl10/teambition/model"

// TBTask 认证接口
type TBTask interface {

	// GetTask 获取任务
	GetTask(taskID string, detailType ...string) (*model.Task, error)

	// AddTask 新增任务
	AddTask(newTask *model.NewTask) (*model.Task, error)

	// DeleteTask 删除任务
	DeleteTask(taskID string) error
}
