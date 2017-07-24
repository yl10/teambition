package services

import (
	"encoding/json"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultOrgService 默认 TBOrganization 接口实现
type DefaultOrgService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultOrgService 创建实例
func NewDefaultOrgService(cli *util.HttpClient, apiURL string) (*DefaultOrgService, error) {
	org := &DefaultOrgService{cli: cli, apiURL: apiURL}
	return org, nil
}

// GetOrg 获取组织
func (t *DefaultOrgService) GetOrg(orgID string) (*model.Organization, error) {
	if orgID == "" {
		return nil, util.NewError("获取组织失败, 没有指定ID")
	}

	reqURL := t.apiURL + "api/organizations/" + orgID

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	var org model.Organization
	err = json.Unmarshal(res, &org)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &org, nil
}

// GetOrgs 获取组织列表
func (t *DefaultOrgService) GetOrgs(withCounts ...bool) ([]model.Organization, error) {
	reqURL := t.apiURL + "api/organizations"
	if len(withCounts) > 0 && withCounts[0] {
		reqURL = reqURL + "?withCounts=true"
	}

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	var os []model.Organization
	err = json.Unmarshal(res, &os)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return os, nil
}

// GetMembers 获取成员
func (t *DefaultOrgService) GetMembers(orgID string, memID ...string) ([]model.Member, error) {
	if orgID == "" {
		return nil, util.NewError("获取成员失败, 未指定组织ID")
	}

	reqURL := t.apiURL + "api/v2/organizations/" + orgID + "/members"
	if len(memID) > 0 && memID[0] != "" {
		reqURL = reqURL + "?_userId=" + memID[0]
	}

	res, err := t.cli.Get(reqURL, nil)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	var ms []model.Member
	err = json.Unmarshal(res, &ms)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return ms, nil
}
