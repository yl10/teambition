package model

import "time"

type App struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Order int    `json:"order"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type CustomField struct {
	Type          string      `json:"type"`
	CustomfieldID string      `json:"_customfieldId"`
	Values        interface{} `json:"values"`
}

type ExecutorOrCreator struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	ID        string `json:"_id"`
}

type LocalLang struct {
	Zh string
	En string
}

type Org struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	Logo        string  `json:"logo"`
	Description string  `json:"description"`
	Plan        OrgPlan `json:"plan"`
	IsPublic    bool    `json:"isPublic"`
	IsExpired   bool    `json:"isExpired"`
}

type OrgPlan struct {
	Status       string    `json:"status"`
	Expired      time.Time `json:"expired"`
	PayType      string    `json:"payType"`
	PaidCount    int       `json:"paidCount"`
	MembersCount int       `json:"membersCount"`
	Days         int       `json:"days"`
	ObjectType   string    `json:"objectType"`
	IsExpired    bool      `json:"isExpired"`
}

type SimpleProject struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type SimpleStage struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type SimpleTask struct {
	ID      string `json:"_id"`
	Content string `json:"content"`
}

type SimpleTasklist struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type SubtaskTotalAndDone struct {
	Total int `json:"total"`
	Done  int `json:"done"`
}

type TaskReminder struct {
	Members []string  `json:"members"`
	Date    time.Time `json:"date"`
	Type    string    `json:"type"`
}

type TBAPIError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
