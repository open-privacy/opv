// Code generated by entc, DO NOT EDIT.

package scope

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-privacy/opv/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CustomID applies equality check predicate on the "custom_id" field. It's identical to CustomIDEQ.
func CustomID(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomID), v))
	})
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CustomIDEQ applies the EQ predicate on the "custom_id" field.
func CustomIDEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomID), v))
	})
}

// CustomIDNEQ applies the NEQ predicate on the "custom_id" field.
func CustomIDNEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomID), v))
	})
}

// CustomIDIn applies the In predicate on the "custom_id" field.
func CustomIDIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCustomID), v...))
	})
}

// CustomIDNotIn applies the NotIn predicate on the "custom_id" field.
func CustomIDNotIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCustomID), v...))
	})
}

// CustomIDGT applies the GT predicate on the "custom_id" field.
func CustomIDGT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustomID), v))
	})
}

// CustomIDGTE applies the GTE predicate on the "custom_id" field.
func CustomIDGTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustomID), v))
	})
}

// CustomIDLT applies the LT predicate on the "custom_id" field.
func CustomIDLT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustomID), v))
	})
}

// CustomIDLTE applies the LTE predicate on the "custom_id" field.
func CustomIDLTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustomID), v))
	})
}

// CustomIDContains applies the Contains predicate on the "custom_id" field.
func CustomIDContains(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustomID), v))
	})
}

// CustomIDHasPrefix applies the HasPrefix predicate on the "custom_id" field.
func CustomIDHasPrefix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustomID), v))
	})
}

// CustomIDHasSuffix applies the HasSuffix predicate on the "custom_id" field.
func CustomIDHasSuffix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustomID), v))
	})
}

// CustomIDEqualFold applies the EqualFold predicate on the "custom_id" field.
func CustomIDEqualFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustomID), v))
	})
}

// CustomIDContainsFold applies the ContainsFold predicate on the "custom_id" field.
func CustomIDContainsFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustomID), v))
	})
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNonce), v))
	})
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNonce), v...))
	})
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNonce), v...))
	})
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNonce), v))
	})
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNonce), v))
	})
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNonce), v))
	})
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNonce), v))
	})
}

// NonceContains applies the Contains predicate on the "nonce" field.
func NonceContains(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNonce), v))
	})
}

// NonceHasPrefix applies the HasPrefix predicate on the "nonce" field.
func NonceHasPrefix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNonce), v))
	})
}

// NonceHasSuffix applies the HasSuffix predicate on the "nonce" field.
func NonceHasSuffix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNonce), v))
	})
}

// NonceEqualFold applies the EqualFold predicate on the "nonce" field.
func NonceEqualFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNonce), v))
	})
}

// NonceContainsFold applies the ContainsFold predicate on the "nonce" field.
func NonceContainsFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNonce), v))
	})
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	})
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDomain), v...))
	})
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.Scope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDomain), v...))
	})
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	})
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	})
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	})
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	})
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	})
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	})
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	})
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	})
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	})
}

// HasFacts applies the HasEdge predicate on the "facts" edge.
func HasFacts() predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FactsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FactsTable, FactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFactsWith applies the HasEdge predicate on the "facts" edge with a given conditions (other predicates).
func HasFactsWith(preds ...predicate.Fact) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FactsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FactsTable, FactsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Scope) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Scope) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Scope) predicate.Scope {
	return predicate.Scope(func(s *sql.Selector) {
		p(s.Not())
	})
}
