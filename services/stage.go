package services

import (
	"encoding/json"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultStageService 默认 TBStage 接口实现
type DefaultStageService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultStageService 创建实例
func NewDefaultStageService(cli *util.HttpClient, apiURL string) (*DefaultStageService, error) {
	stage := &DefaultStageService{cli: cli, apiURL: apiURL}
	return stage, nil
}

// GetStage 获取任务列表
func (t *DefaultStageService) GetStage(stageID string) (*model.Stage, error) {
	if stageID == "" {
		return nil, util.NewError("查询任务列表失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/Stages/" + stageID

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, err
	}

	var stage model.Stage
	if err = json.Unmarshal(res, &stage); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &stage, nil
}

// AddStage 新增Stage
func (t *DefaultStageService) AddStage(newStage *model.NewStage) (*model.Stage, error) {
	reqURL := t.apiURL + "api/stages"
	res, err := t.cli.Post(reqURL, newStage)
	if err != nil {
		return nil, err
	}

	var stage model.Stage
	if err = json.Unmarshal(res, &stage); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &stage, nil
}

// DeleteStage 删除Stage
func (t *DefaultStageService) DeleteStage(stageID string) error {
	if stageID == "" {
		return util.NewError("删除Stage失败, 没有StageID")
	}

	reqURL := t.apiURL + "api/stages/" + stageID

	_, err := t.cli.Delete(reqURL, nil)
	return err
}
