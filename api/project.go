package api

import "github.com/yl10/teambition/model"

// TBProject 认证接口
type TBProject interface {
	GetProject(projectID string, teamID string, isArchived bool, isWithActivity bool) (*model.Project, error)

	// AddProject 新增项目
	AddProject(newProject *model.NewProject) (*model.Project, error)

	// DeleteProject 删除项目
	DeleteProject(projectID string) error
}
