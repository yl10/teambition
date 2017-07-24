package model

type Stage struct {
	ID         string `json:"_id"`
	ProjectID  string `json:"_projectId"`
	TasklistID string `json:"_tasklistId"`
	Name       string `json:"name"`
	Order      int    `json:"order"`
	TotalCount int    `json:"totalCount"`
	IsArchived bool   `json:"isArchived"`
}

type NewStage struct {
	Name       string `json:"name"`
	PrevID     string `json:"_prevId"`
	TasklistID string `json:"_tasklistId"`
}
