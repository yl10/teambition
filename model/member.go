package model

import "time"

type Member struct {
	ID                   string      `json:"_id"`
	UserID               string      `json:"_userId"`
	BoundToObjectID      string      `json:"_boundToObjectId"`
	BoundToObjectType    string      `json:"boundToObjectType"`
	RoleID               int         `json:"_roleId"`
	Visited              time.Time   `json:"visited"`
	Joined               time.Time   `json:"joined"`
	PushStatus           bool        `json:"pushStatus"`
	Plan                 OrgPlan     `json:"plan"`
	Teams                []string    `json:"teams"`
	IsDisabled           bool        `json:"isDisabled"`
	IsOrgMember          bool        `json:"isOrgMember"`
	Profile              interface{} `json:"profile"`
	ProjectExperienceIDs []string    `json:"projectExperienceIds"`
	Nickname             string      `json:"nickname"`
	HasVisited           bool        `json:"hasVisited"`
	MemberID             string      `json:"_memberId"`
	Location             string      `json:"location"`
	Website              string      `json:"website"`
	LatestActived        time.Time   `json:"latestActived"`
	IsActive             bool        `json:"isActive"`
	Email                string      `json:"email"`
	Name                 string      `json:"name"`
	AvatarURL            string      `json:"avatarUrl"`
	Title                string      `json:"title"`
	IsBlock              bool        `json:"isBlock"`
	Pinyin               string      `json:"pinyin"`
	Py                   string      `json:"py"`
	ProjectExperience    []string    `json:"projectExperience"`
}
