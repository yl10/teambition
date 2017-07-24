package model

import "time"

type Application struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
	Status      string    `json:"status"`
	Description LocalLang `json:"description"`
	Title       LocalLang `json:"title"`
	Type        int       `json:"type"`
}
