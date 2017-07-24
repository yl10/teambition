package model

import "time"

const (
	TASK_PRIORITY_NORMAL = iota
	TASK_PRIORITY_IMPORTANT
	TASK_PRIORITY_URGENT
)

type Task struct {
	ID               string              `json:"_id"`
	Content          string              `json:"content"`
	Note             string              `json:"note"`
	Accomplished     string              `json:"accomplished"`
	StartDate        time.Time           `json:"startDate"`
	DueDate          time.Time           `json:"dueDate"`
	Priority         int                 `json:"priority"`
	IsDone           bool                `json:"isDone"`
	IsArchived       bool                `json:"isArchived"`
	IsDeleted        bool                `json:"isDeleted"`
	Created          time.Time           `json:"created"`
	Updated          time.Time           `json:"updated"`
	Visible          string              `json:"visible"`
	StageID          string              `json:"_stageId"`
	CreatorID        string              `json:"_creatorId"`
	TasklistID       string              `json:"_tasklistId"`
	ProjectID        string              `json:"_projectId"`
	ExecutorID       string              `json:"_executorId"`
	InvolveMembers   []string            `json:"involveMembers"`
	TagIDs           []string            `json:"tagIds"`
	Recurrence       string              `json:"recurrence"`
	Pos              int                 `json:"pos"`
	SourceID         string              `json:"_sourceId"`
	SourceDate       string              `json:"sourceDate"`
	Subtasks         []SubTask           `json:"subtasks"`
	SubtaskIDs       []string            `json:"subtaskIds"`
	Source           string              `json:"source"`
	Customfields     []CustomField       `json:"customfields"`
	Involvers        []ExecutorOrCreator `json:"involvers"`
	CommentsCount    int                 `json:"commentsCount"`
	AttachmentsCount int                 `json:"attachmentsCount"`
	LikesCount       int                 `json:"likesCount"`
	ObjectlinksCount int                 `json:"objectlinksCount"`
	ShareStatus      int                 `json:"shareStatus"`
	Reminder         TaskReminder        `json:"reminder"`
	SubtaskCount     SubtaskTotalAndDone `json:"subtaskCount"`
	Executor         ExecutorOrCreator   `json:"executor"`
	Stage            SimpleStage         `json:"stage"`
	Tasklist         SimpleTasklist      `json:"tasklist"`
	Type             string              `json:"type"`
	IsFavorite       bool                `json:"isFavorite"`
	Project          SimpleProject       `json:"project"`
	UniqueID         int                 `json:"uniqueId"`
	URL              string              `json:"url"`
}

type NewTask struct {
	Content        string     `json:"content"`
	TasklistID     string     `json:"_tasklistId"`
	StageID        string     `json:"_stageId,omitempty"`
	ExecutorID     string     `json:"_executorId,omitempty"`
	InvolveMembers []string   `json:"involveMembers,omitempty"`
	DueDate        *time.Time `json:"dueDate,omitempty"`
	Priority       int        `json:"priority,omitempty"`
	Recurrence     []string   `json:"recurrence,omitempty"`
	TagIds         []string   `json:"tagIds,omitempty"`
	Note           string     `json:"note,omitempty"`
}
