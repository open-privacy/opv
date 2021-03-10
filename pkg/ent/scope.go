// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-privacy/opv/pkg/ent/scope"
)

// Scope is the model entity for the Scope schema.
type Scope struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// CustomID holds the value of the "custom_id" field.
	CustomID string `json:"custom_id,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"-"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScopeQuery when eager-loading is set.
	Edges ScopeEdges `json:"edges"`
}

// ScopeEdges holds the relations/edges for other nodes in the graph.
type ScopeEdges struct {
	// Facts holds the value of the facts edge.
	Facts []*Fact `json:"facts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FactsOrErr returns the Facts value or an error if the edge
// was not loaded in eager-loading.
func (e ScopeEdges) FactsOrErr() ([]*Fact, error) {
	if e.loadedTypes[0] {
		return e.Facts, nil
	}
	return nil, &NotLoadedError{edge: "facts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Scope) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case scope.FieldID, scope.FieldCustomID, scope.FieldNonce, scope.FieldDomain:
			values[i] = &sql.NullString{}
		case scope.FieldCreatedAt, scope.FieldUpdatedAt, scope.FieldDeletedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Scope", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Scope fields.
func (s *Scope) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scope.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case scope.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case scope.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case scope.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case scope.FieldCustomID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom_id", values[i])
			} else if value.Valid {
				s.CustomID = value.String
			}
		case scope.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				s.Nonce = value.String
			}
		case scope.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		}
	}
	return nil
}

// QueryFacts queries the "facts" edge of the Scope entity.
func (s *Scope) QueryFacts() *FactQuery {
	return (&ScopeClient{config: s.config}).QueryFacts(s)
}

// Update returns a builder for updating this Scope.
// Note that you need to call Scope.Unwrap() before calling this method if this Scope
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Scope) Update() *ScopeUpdateOne {
	return (&ScopeClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Scope entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Scope) Unwrap() *Scope {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Scope is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Scope) String() string {
	var builder strings.Builder
	builder.WriteString("Scope(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", custom_id=")
	builder.WriteString(s.CustomID)
	builder.WriteString(", nonce=<sensitive>")
	builder.WriteString(", domain=")
	builder.WriteString(s.Domain)
	builder.WriteByte(')')
	return builder.String()
}

// Scopes is a parsable slice of Scope.
type Scopes []*Scope

func (s Scopes) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
