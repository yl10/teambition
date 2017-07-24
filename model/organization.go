package model

import (
	"time"
)

type Organization struct {
	ID            string        `json:"_id"`
	CreatorID     string        `json:"_creatorId"`
	Category      string        `json:"category"`
	Name          string        `json:"name"`
	Logo          string        `json:"logo"`
	Location      string        `json:"location"`
	Description   string        `json:"description"`
	Background    string        `json:"background"`
	Website       string        `json:"website"`
	ProjectIds    []string      `json:"projectIds"`
	Plan          OrgPlan       `json:"plan"`
	Pinyin        string        `json:"pinyin"`
	Py            string        `json:"py"`
	Dividers      []interface{} `json:"dividers"`
	IsPublic      bool          `json:"isPublic"`
	RoleID        int           `json:"_roleId"`
	DefaultRoleID string        `json:"_defaultRoleId"`
	DefaultTeamID string        `json:"_defaultTeamId"`
	StaffTypes    []string      `json:"staffTypes"`
	Positions     []string      `json:"positions"`
	HasRight      int           `json:"hasRight"`
	Created       time.Time     `json:"created"`
	Updated       time.Time     `json:"updated"`
}
