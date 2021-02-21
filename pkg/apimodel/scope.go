package apimodel

import (
	"time"
)

type CreateScope struct {
	CustomID string `json:"type"`
}

type Scope struct {
	ID       string `json:"id"`
	CustomID string `json:"custom_id"`

	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
