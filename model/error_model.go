package model

import "time"

type Error struct {
	CreatedAt  time.Time `json:"createdAt"`
	App        string    `json:"app"`
	Class      string    `json:"class"`
	Method     string    `json:"method"`
	Message    string    `json:"message"`
	StackTrace string    `json:"stackTrace"`
	ExtraInfo  string    `json:"extraInfo"`
}
