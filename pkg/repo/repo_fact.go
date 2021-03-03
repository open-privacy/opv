package repo

import (
	"errors"
	"fmt"

	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/upper/db/v4"
)

// Fact is the repo pattern of model
type Fact struct {
	Base

	// Fields
	Domain         string `db:"domain,omitempty"`
	HashedValue    string `db:"hashed_value,omitempty"`
	EncryptedValue string `db:"encrypted_value,omitempty"`

	// Associations
	ScopeCustomID string `db:"scope_custom_id,omitempty"`
	FactTypeSlug  string `db:"fact_type_slug,omitempty"`

	// Helpers
	encryptor crypto.Encryptor
	hasher    crypto.Hasher
}

// Store implements the db.Record interface
func (f *Fact) Store(sess db.Session) db.Store {
	return sess.Collection("facts")
}

// WithEncrytor returns a fact with encryptor set
func (f *Fact) WithEncrytor(encryptor crypto.Encryptor) *Fact {
	f.encryptor = encryptor
	return f
}

// WithHasher returns a fact with hasher set
func (f *Fact) WithHasher(hasher crypto.Hasher) *Fact {
	f.hasher = hasher
	return f
}

// Insert creates a new fact to the database
func (f *Fact) Insert(sess db.Session, value string) error {
	if f.encryptor == nil || f.hasher == nil {
		return fmt.Errorf("encryptor or hasher cannot be nil")
	}
	if s, err := f.createScopeIfNotExists(sess); err != nil {
		return err
	}
	if ft, err := f.createFactTypeIfNotExists(sess); err != nil {
		return err
	}

	encryptedValue, err := f.encryptor.Encrypt()

	return nil
}

func (f *Fact) createScopeIfNotExists(sess db.Session) (*Scope, error) {
	s := &Scope{
		CustomID: f.ScopeCustomID,
		Domain:   f.Domain,
	}
	sCollection := s.Store(sess)

	err := sCollection.Find(db.Cond{"custom_id": s.CustomID}).One(s)
	if errors.Is(err, db.ErrNoMoreRows) {
		if _, err := sCollection.Insert(s); err != nil {
			return nil, err
		}
		return s, nil
	}
	return s, err
}

func (f *Fact) createFactTypeIfNotExists(sess db.Session) (*FactType, error) {
	ft := &FactType{
		Slug:           f.FactTypeSlug,
		BuiltIn:        false,
		ValidationRule: "",
	}
	ftCollection := ft.Store(sess)

	err := ftCollection.Find(db.Cond{"slug": ft.Slug}).One(ft)
	if errors.Is(err, db.ErrNoMoreRows) {
		if _, err := ftCollection.Insert(ft); err != nil {
			return nil, err
		}
		return ft, nil
	}
	return ft, err
}
