package services

import (
	"encoding/json"

	"net/url"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultProjectService 默认 TBProject 接口实现
type DefaultProjectService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultProjectService 创建实例
func NewDefaultProjectService(cli *util.HttpClient, apiURL string) (*DefaultProjectService, error) {
	project := &DefaultProjectService{cli: cli, apiURL: apiURL}
	return project, nil
}

// GetProject 获取项目
func (t *DefaultProjectService) GetProject(projectID string, teamID string, isArchived bool, isWithActivity bool) (*model.Project, error) {
	if projectID == "" {
		return nil, util.NewError("查询项目失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/projects/" + projectID
	params := make(url.Values)
	if teamID != "" {
		params.Add("_teamId", teamID)
	}

	if isArchived {
		params.Add("isArchived", "true")
	}
	if isWithActivity {
		params.Add("isWithActivity", "true")
	}

	res, err := t.cli.Get(reqURL, params.Encode())
	if err != nil {
		return nil, err
	}

	var project model.Project
	if err = json.Unmarshal(res, &project); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &project, nil
}

// AddProject 新增项目
func (t *DefaultProjectService) AddProject(newProject *model.NewProject) (*model.Project, error) {
	reqURL := t.apiURL + "api/projects"
	res, err := t.cli.Post(reqURL, newProject)
	if err != nil {
		return nil, err
	}

	var project model.Project
	if err = json.Unmarshal(res, &project); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &project, nil
}

// DeleteProject 删除项目
func (t *DefaultProjectService) DeleteProject(projectID string) error {
	if projectID == "" {
		return util.NewError("删除项目失败, 没有项目ID")
	}

	reqURL := t.apiURL + "api/projects/" + projectID

	_, err := t.cli.Delete(reqURL, nil)
	return err
}
