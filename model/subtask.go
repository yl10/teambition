package model

import "time"

type SubTask struct {
	ID         string            `json:"_id"`
	ProjectID  string            `json:"_projectId"`
	CreatorID  string            `json:"_creatorId"`
	Created    time.Time         `json:"created"`
	Content    string            `json:"content"`
	IsDone     bool              `json:"isDone"`
	ExecutorID string            `json:"_executorId"`
	TaskID     string            `json:"_taskId"`
	DueDate    time.Time         `json:"dueDate"`
	Order      int               `json:"order"`
	Executor   ExecutorOrCreator `json:"executor"`
	Updated    time.Time         `json:"updated"`
	Type       string            `json:"type"`
	Project    SimpleProject     `json:"project"`
	Task       SimpleTask        `json:"task"`
}
