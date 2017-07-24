package model

import "time"

type Tasklist struct {
	ID           string    `json:"_id"`
	Title        string    `json:"title"`
	ProjectID    string    `json:"_projectId"`
	CreatorID    string    `json:"_creatorId"`
	Description  string    `json:"description"`
	IsArchived   bool      `json:"isArchived"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	StageIDs     []string  `json:"stageIds"`
	DoneCount    int       `json:"doneCount"`
	UndoneCount  int       `json:"undoneCount"`
	ExpiredCount int       `json:"expiredCount"`
	RecentCount  int       `json:"recentCount"`
	TotalCount   int       `json:"totalCount"`
	HasStages    []Stage   `json:"hasStages"`
	Lock         string    `json:"lock"`
	SortMethod   string    `json:"sortMethod"`
}

type NewTasklist struct {
	Title       string `json:"title"`
	ProjectID   string `json:"_projectId"`
	Description string `json:"description"`
	TemplateID  string `json:"_templateId"`
}
