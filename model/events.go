package model

import (
	"time"
)

type Events struct {
	ID               string              `json:"_id"`
	EndDate          time.Time           `json:"endDate"`
	StartDate        time.Time           `json:"startDate"`
	ProjectID        string              `json:"_projectId"`
	Location         string              `json:"location"`
	Content          string              `json:"content"`
	Title            string              `json:"title"`
	CreatorID        string              `json:"_creatorId"`
	TagIDs           []string            `json:"tagIds"`
	Updated          time.Time           `json:"updated"`
	Created          time.Time           `json:"created"`
	Visible          string              `json:"visible"`
	IsArchived       bool                `json:"isArchived"`
	InvolveMembers   []string            `json:"involveMembers"`
	Status           string              `json:"status"`
	UntilDate        time.Time           `json:"untilDate"`
	SourceID         string              `json:"_sourceId"`
	SourceDate       time.Time           `json:"sourceDate"`
	Recurrence       []string            `json:"recurrence"`
	Reminders        []string            `json:"reminders"`
	Involvers        []ExecutorOrCreator `json:"involvers"`
	Creator          ExecutorOrCreator   `json:"creator"`
	CommentsCount    int                 `json:"commentsCount"`
	AttachmentsCount int                 `json:"attachmentsCount"`
	LikesCount       int                 `json:"likesCount"`
	ObjectlinksCount int                 `json:"objectlinksCount"`
	IsFavorite       bool                `json:"isFavorite"`
	ShareStatus      int                 `json:"shareStatus"`
}

type NewEvents struct {
	ProjectID        string              `json:"_projectId"`
	Title            string              `json:"title"`
	EndDate          time.Time           `json:"endDate"`
	StartDate        time.Time           `json:"startDate"`
	Location         string              `json:"location"`
	Content          string              `json:"content"`
	Recurrence       []string            `json:"recurrence"`
	Reminders        []string            `json:"reminders"`
	TagIDs           []string            `json:"tagIds"`
	Visible          string              `json:"visible"`
	InvolveMembers   []string            `json:"involveMembers"`
	SourceID         string              `json:"_sourceId"`
	SourceDate       time.Time           `json:"sourceDate"`
}
