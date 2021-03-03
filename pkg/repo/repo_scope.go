package repo

import (
	"github.com/dchest/uniuri"
	"github.com/upper/db/v4"
)

// Scope is the repo pattern of Scope struct
type Scope struct {
	Base

	CustomID string `db:"custom_id,omitempty"`
	Nonce    string `db:"nonce,omitempty"`
	Domain   string `db:"domain,omitempty"`
}

// Store implements the db.Record interface
func (s *Scope) Store(sess db.Session) db.Store {
	return sess.Collection("scopes")
}

// 
func CreateScope(sess db.Session, domain string, customID string) (*Scope, error) {
	err := s.Store().Find(db.Cond({"custom_id": customID})).One(s)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {

			err := s.Store(sess).InsertReturning(s)

			// If transation is not supported, e.g. mongodb, insert first and then get
			if errors.Is(err, db.ErrUnsupported) {
				err := s.Store(sess).Insert()
			}
		}
		return err
	}
}