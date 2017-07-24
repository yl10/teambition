package model

import "time"

type Project struct {
	CreatorID           string            `json:"_creatorId"`
	DefaultCollectionID string            `json:"_defaultCollectionId"`
	DefaultRoleID       string            `json:"_defaultRoleId"`
	ID                  string            `json:"_id"`
	OrgRoleID           int               `json:"_orgRoleId"`
	OrganizationID      string            `json:"_organizationId"`
	RoleID              int               `json:"_roleId"`
	RootCollectionID    string            `json:"_rootCollectionId"`
	Applications        []App             `json:"applications"`
	Category            string            `json:"category"`
	Created             time.Time         `json:"created"`
	Creator             ExecutorOrCreator `json:"creator"`
	CustomFields        []CustomField     `json:"customFields"`
	Description         string            `json:"description"`
	EventsCount         int               `json:"eventsCount"`
	HasOrgRight         int               `json:"hasOrgRight"`
	HasRight            int               `json:"hasRight"`
	InviteLink          string            `json:"inviteLink"`
	IsArchived          bool              `json:"isArchived"`
	IsPublic            bool              `json:"isPublic"`
	IsStar              bool              `json:"isStar"`
	Logo                string            `json:"logo"`
	MembersCount        int               `json:"membersCount"`
	Name                string            `json:"name"`
	Organization        Org               `json:"organization"`
	Pinyin              string            `json:"pinyin"`
	PostsCount          int               `json:"postsCount"`
	PushStatus          bool              `json:"pushStatus"`
	Py                  string            `json:"py"`
	ShortLink           string            `json:"shortLink"`
	SortMethod          string            `json:"sortMethod"`
	StarsCount          int               `json:"starsCount"`
	TagsCount           int               `json:"tagsCount"`
	TasksCount          int               `json:"tasksCount"`
	UniqueIDPrefix      string            `json:"uniqueIdPrefix"`
	UnreadCount         int               `json:"unreadCount"`
	Updated             time.Time         `json:"updated"`
	Visibility          string            `json:"visibility"`
	WorksCount          int               `json:"worksCount"`
}

type NewProject struct {
	Name           string `json:"name"`
	OrganizationId string `json:"_organizationId"`
	Description    string `json:"description"`
	Logo           string `json:"logo"`
	Category       string `json:"category"`
	DividerIndex   int    `json:"dividerIndex"`
	Visibility     string `json:"visibility"`
}
