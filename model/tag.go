package model

import "time"

type Tag struct {
	ID          string    `json:"_id"`
	ProjectID   string    `json:"_projectId"`
	CreatorID   string    `json:"_creatorId"`
	Name        string    `json:"name"`
	IsArchived  bool      `json:"isArchived"`
	Color       string    `json:"color"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	PostsCount  int       `json:"postsCount"`
	TasksCount  int       `json:"tasksCount"`
	EventsCount int       `json:"eventsCount"`
	WorksCount  int       `json:"worksCount"`
}

type NewTag struct {
	Name      string `json:"name"`
	ProjectID string `json:"_projectId"`
}
