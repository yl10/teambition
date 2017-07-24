package api

import "github.com/yl10/teambition/model"

// TBStage 认证接口
type TBStage interface {

	// GetStage 获取Stage
	GetStage(stageID string) (*model.Stage, error)

	// AddStage 新增Stage
	AddStage(newStage *model.NewStage) (*model.Stage, error)

	// DeleteStage 删除Stage
	DeleteStage(StageID string) error
}
