package model

import "time"

type Client struct {
	CreatedAt time.Time `json:"createdAt"`
	App       string    `json:"app"`
	Type      int8      `json:"type"`
	Message   string    `json:"message"`
	ExtraInfo string    `json:"extraInfo"`
}
