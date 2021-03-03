package repo

import (
	"time"

	"github.com/dchest/uniuri"
	"github.com/upper/db/v4"
)

// Base is the base struct for repo pattern of models
type Base struct {
	ID        string     `db:"id,omitempty"`
	CreatedAt time.Time  `db:"created_at,omitempty"`
	UpdatedAt time.Time  `db:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at,omitempty"`
}

// BeforeCreate sets default ID and CreatedAt
func (b *Base) BeforeCreate(sess db.Session) error {
	b.ID = uniuri.NewLen(uniuri.UUIDLen)
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
	return nil
}

// BeforeUpdate sets default UpdatedAt
func (b *Base) BeforeUpdate(sess db.Session) error {
	b.UpdatedAt = time.Now()
	return nil
}
