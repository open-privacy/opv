package apimodel

import (
	"time"
)

type CreateScope struct {
	Type      string     `json:"type"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type Scope struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
}
