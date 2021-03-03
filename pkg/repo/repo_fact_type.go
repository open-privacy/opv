package repo

import "github.com/upper/db/v4"

// FactType is the repo pattern of FactType struct
type FactType struct {
	Base

	Slug           string `db:"slug" validate:"required"`
	BuiltIn        bool   `db:"built_in"`
	ValidationRule string `db:"validation_rule"`
}

// Store implements the db.Record interface
func (ft *FactType) Store(sess db.Session) db.Store {
	return sess.Collection("fact_types")
}
