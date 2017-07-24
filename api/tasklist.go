package api

import "github.com/yl10/teambition/model"

// TBTasklist 认证接口
type TBTasklist interface {
	// GetTasklist 获取分组
	GetTasklist(tasklistID string) (*model.Tasklist, error)

	// AddTasklist 新增分组
	AddTasklist(newTasklist *model.NewTasklist) (*model.Tasklist, error)

	// DeleteTasklist 删除分组
	DeleteTasklist(tasklistID string) error
}
