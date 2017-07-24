package api

import "github.com/yl10/teambition/model"

// TBOrganization 认证接口
type TBOrganization interface {
	// GetOrg 获取组织
	GetOrg(orgID string) (*model.Organization, error)

	// GetOrgs 获取组织列表
	GetOrgs(withCounts ...bool) ([]model.Organization, error)

	// GetMembers 获取成员
	GetMembers(orgID string, memID ...string) ([]model.Member, error)
}
