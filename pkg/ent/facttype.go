// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/open-privacy/opv/pkg/ent/facttype"
)

// FactType is the model entity for the FactType schema.
type FactType struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Builtin holds the value of the "builtin" field.
	Builtin bool `json:"builtin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FactTypeQuery when eager-loading is set.
	Edges FactTypeEdges `json:"edges"`
}

// FactTypeEdges holds the relations/edges for other nodes in the graph.
type FactTypeEdges struct {
	// Facts holds the value of the facts edge.
	Facts []*Fact `json:"facts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FactsOrErr returns the Facts value or an error if the edge
// was not loaded in eager-loading.
func (e FactTypeEdges) FactsOrErr() ([]*Fact, error) {
	if e.loadedTypes[0] {
		return e.Facts, nil
	}
	return nil, &NotLoadedError{edge: "facts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FactType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case facttype.FieldBuiltin:
			values[i] = &sql.NullBool{}
		case facttype.FieldSlug:
			values[i] = &sql.NullString{}
		case facttype.FieldCreateTime, facttype.FieldUpdateTime:
			values[i] = &sql.NullTime{}
		case facttype.FieldID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FactType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FactType fields.
func (ft *FactType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case facttype.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ft.ID = *value
			}
		case facttype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ft.CreateTime = value.Time
			}
		case facttype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ft.UpdateTime = value.Time
			}
		case facttype.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				ft.Slug = value.String
			}
		case facttype.FieldBuiltin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field builtin", values[i])
			} else if value.Valid {
				ft.Builtin = value.Bool
			}
		}
	}
	return nil
}

// QueryFacts queries the "facts" edge of the FactType entity.
func (ft *FactType) QueryFacts() *FactQuery {
	return (&FactTypeClient{config: ft.config}).QueryFacts(ft)
}

// Update returns a builder for updating this FactType.
// Note that you need to call FactType.Unwrap() before calling this method if this FactType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FactType) Update() *FactTypeUpdateOne {
	return (&FactTypeClient{config: ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the FactType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FactType) Unwrap() *FactType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FactType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FactType) String() string {
	var builder strings.Builder
	builder.WriteString("FactType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ft.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ft.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", slug=")
	builder.WriteString(ft.Slug)
	builder.WriteString(", builtin=")
	builder.WriteString(fmt.Sprintf("%v", ft.Builtin))
	builder.WriteByte(')')
	return builder.String()
}

// FactTypes is a parsable slice of FactType.
type FactTypes []*FactType

func (ft FactTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
