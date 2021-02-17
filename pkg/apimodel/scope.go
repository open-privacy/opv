package apimodel

import (
	"time"

	"github.com/google/uuid"
)

type Scope struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Type  string    `json:"type,omitempty"`
	Nonce uuid.UUID `json:"nonce,omitempty"`

	CreateTime time.Time  `json:"create_time,omitempty"`
	UpdateTime time.Time  `json:"update_time,omitempty"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
}
