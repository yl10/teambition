package api

import "github.com/yl10/teambition/model"

// TBEvents 认证接口
type TBEvents interface {
	// AddEvents 新增分组
	AddEvents(newEvents *model.NewEvents) (*model.Events, error)

	// DeleteEvents 删除分组
	DeleteEvents(eventsID string) error
}
