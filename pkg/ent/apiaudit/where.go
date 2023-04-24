// Code generated by entc, DO NOT EDIT.

package apiaudit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/roney492/opv/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
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
func IDGT(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Plane applies equality check predicate on the "plane" field. It's identical to PlaneEQ.
func Plane(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlane), v))
	})
}

// HashedGrantToken applies equality check predicate on the "hashed_grant_token" field. It's identical to HashedGrantTokenEQ.
func HashedGrantToken(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedGrantToken), v))
	})
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// HTTPPath applies equality check predicate on the "http_path" field. It's identical to HTTPPathEQ.
func HTTPPath(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTTPPath), v))
	})
}

// HTTPMethod applies equality check predicate on the "http_method" field. It's identical to HTTPMethodEQ.
func HTTPMethod(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTTPMethod), v))
	})
}

// SentHTTPStatus applies equality check predicate on the "sent_http_status" field. It's identical to SentHTTPStatusEQ.
func SentHTTPStatus(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentHTTPStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// PlaneEQ applies the EQ predicate on the "plane" field.
func PlaneEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlane), v))
	})
}

// PlaneNEQ applies the NEQ predicate on the "plane" field.
func PlaneNEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlane), v))
	})
}

// PlaneIn applies the In predicate on the "plane" field.
func PlaneIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlane), v...))
	})
}

// PlaneNotIn applies the NotIn predicate on the "plane" field.
func PlaneNotIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlane), v...))
	})
}

// PlaneGT applies the GT predicate on the "plane" field.
func PlaneGT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlane), v))
	})
}

// PlaneGTE applies the GTE predicate on the "plane" field.
func PlaneGTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlane), v))
	})
}

// PlaneLT applies the LT predicate on the "plane" field.
func PlaneLT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlane), v))
	})
}

// PlaneLTE applies the LTE predicate on the "plane" field.
func PlaneLTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlane), v))
	})
}

// PlaneContains applies the Contains predicate on the "plane" field.
func PlaneContains(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlane), v))
	})
}

// PlaneHasPrefix applies the HasPrefix predicate on the "plane" field.
func PlaneHasPrefix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlane), v))
	})
}

// PlaneHasSuffix applies the HasSuffix predicate on the "plane" field.
func PlaneHasSuffix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlane), v))
	})
}

// PlaneEqualFold applies the EqualFold predicate on the "plane" field.
func PlaneEqualFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlane), v))
	})
}

// PlaneContainsFold applies the ContainsFold predicate on the "plane" field.
func PlaneContainsFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlane), v))
	})
}

// HashedGrantTokenEQ applies the EQ predicate on the "hashed_grant_token" field.
func HashedGrantTokenEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenNEQ applies the NEQ predicate on the "hashed_grant_token" field.
func HashedGrantTokenNEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenIn applies the In predicate on the "hashed_grant_token" field.
func HashedGrantTokenIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHashedGrantToken), v...))
	})
}

// HashedGrantTokenNotIn applies the NotIn predicate on the "hashed_grant_token" field.
func HashedGrantTokenNotIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHashedGrantToken), v...))
	})
}

// HashedGrantTokenGT applies the GT predicate on the "hashed_grant_token" field.
func HashedGrantTokenGT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenGTE applies the GTE predicate on the "hashed_grant_token" field.
func HashedGrantTokenGTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenLT applies the LT predicate on the "hashed_grant_token" field.
func HashedGrantTokenLT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenLTE applies the LTE predicate on the "hashed_grant_token" field.
func HashedGrantTokenLTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenContains applies the Contains predicate on the "hashed_grant_token" field.
func HashedGrantTokenContains(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenHasPrefix applies the HasPrefix predicate on the "hashed_grant_token" field.
func HashedGrantTokenHasPrefix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenHasSuffix applies the HasSuffix predicate on the "hashed_grant_token" field.
func HashedGrantTokenHasSuffix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenIsNil applies the IsNil predicate on the "hashed_grant_token" field.
func HashedGrantTokenIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHashedGrantToken)))
	})
}

// HashedGrantTokenNotNil applies the NotNil predicate on the "hashed_grant_token" field.
func HashedGrantTokenNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHashedGrantToken)))
	})
}

// HashedGrantTokenEqualFold applies the EqualFold predicate on the "hashed_grant_token" field.
func HashedGrantTokenEqualFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHashedGrantToken), v))
	})
}

// HashedGrantTokenContainsFold applies the ContainsFold predicate on the "hashed_grant_token" field.
func HashedGrantTokenContainsFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHashedGrantToken), v))
	})
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	})
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func DomainNotIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
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
func DomainGT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	})
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	})
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	})
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	})
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	})
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	})
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	})
}

// DomainIsNil applies the IsNil predicate on the "domain" field.
func DomainIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDomain)))
	})
}

// DomainNotNil applies the NotNil predicate on the "domain" field.
func DomainNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDomain)))
	})
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	})
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	})
}

// HTTPPathEQ applies the EQ predicate on the "http_path" field.
func HTTPPathEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathNEQ applies the NEQ predicate on the "http_path" field.
func HTTPPathNEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathIn applies the In predicate on the "http_path" field.
func HTTPPathIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHTTPPath), v...))
	})
}

// HTTPPathNotIn applies the NotIn predicate on the "http_path" field.
func HTTPPathNotIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHTTPPath), v...))
	})
}

// HTTPPathGT applies the GT predicate on the "http_path" field.
func HTTPPathGT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathGTE applies the GTE predicate on the "http_path" field.
func HTTPPathGTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathLT applies the LT predicate on the "http_path" field.
func HTTPPathLT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathLTE applies the LTE predicate on the "http_path" field.
func HTTPPathLTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathContains applies the Contains predicate on the "http_path" field.
func HTTPPathContains(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathHasPrefix applies the HasPrefix predicate on the "http_path" field.
func HTTPPathHasPrefix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathHasSuffix applies the HasSuffix predicate on the "http_path" field.
func HTTPPathHasSuffix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathIsNil applies the IsNil predicate on the "http_path" field.
func HTTPPathIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHTTPPath)))
	})
}

// HTTPPathNotNil applies the NotNil predicate on the "http_path" field.
func HTTPPathNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHTTPPath)))
	})
}

// HTTPPathEqualFold applies the EqualFold predicate on the "http_path" field.
func HTTPPathEqualFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHTTPPath), v))
	})
}

// HTTPPathContainsFold applies the ContainsFold predicate on the "http_path" field.
func HTTPPathContainsFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHTTPPath), v))
	})
}

// HTTPMethodEQ applies the EQ predicate on the "http_method" field.
func HTTPMethodEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodNEQ applies the NEQ predicate on the "http_method" field.
func HTTPMethodNEQ(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodIn applies the In predicate on the "http_method" field.
func HTTPMethodIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHTTPMethod), v...))
	})
}

// HTTPMethodNotIn applies the NotIn predicate on the "http_method" field.
func HTTPMethodNotIn(vs ...string) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHTTPMethod), v...))
	})
}

// HTTPMethodGT applies the GT predicate on the "http_method" field.
func HTTPMethodGT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodGTE applies the GTE predicate on the "http_method" field.
func HTTPMethodGTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodLT applies the LT predicate on the "http_method" field.
func HTTPMethodLT(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodLTE applies the LTE predicate on the "http_method" field.
func HTTPMethodLTE(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodContains applies the Contains predicate on the "http_method" field.
func HTTPMethodContains(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodHasPrefix applies the HasPrefix predicate on the "http_method" field.
func HTTPMethodHasPrefix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodHasSuffix applies the HasSuffix predicate on the "http_method" field.
func HTTPMethodHasSuffix(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodIsNil applies the IsNil predicate on the "http_method" field.
func HTTPMethodIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHTTPMethod)))
	})
}

// HTTPMethodNotNil applies the NotNil predicate on the "http_method" field.
func HTTPMethodNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHTTPMethod)))
	})
}

// HTTPMethodEqualFold applies the EqualFold predicate on the "http_method" field.
func HTTPMethodEqualFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHTTPMethod), v))
	})
}

// HTTPMethodContainsFold applies the ContainsFold predicate on the "http_method" field.
func HTTPMethodContainsFold(v string) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHTTPMethod), v))
	})
}

// SentHTTPStatusEQ applies the EQ predicate on the "sent_http_status" field.
func SentHTTPStatusEQ(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusNEQ applies the NEQ predicate on the "sent_http_status" field.
func SentHTTPStatusNEQ(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusIn applies the In predicate on the "sent_http_status" field.
func SentHTTPStatusIn(vs ...int) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSentHTTPStatus), v...))
	})
}

// SentHTTPStatusNotIn applies the NotIn predicate on the "sent_http_status" field.
func SentHTTPStatusNotIn(vs ...int) predicate.APIAudit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.APIAudit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSentHTTPStatus), v...))
	})
}

// SentHTTPStatusGT applies the GT predicate on the "sent_http_status" field.
func SentHTTPStatusGT(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusGTE applies the GTE predicate on the "sent_http_status" field.
func SentHTTPStatusGTE(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusLT applies the LT predicate on the "sent_http_status" field.
func SentHTTPStatusLT(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusLTE applies the LTE predicate on the "sent_http_status" field.
func SentHTTPStatusLTE(v int) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSentHTTPStatus), v))
	})
}

// SentHTTPStatusIsNil applies the IsNil predicate on the "sent_http_status" field.
func SentHTTPStatusIsNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSentHTTPStatus)))
	})
}

// SentHTTPStatusNotNil applies the NotNil predicate on the "sent_http_status" field.
func SentHTTPStatusNotNil() predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSentHTTPStatus)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.APIAudit) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.APIAudit) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
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
func Not(p predicate.APIAudit) predicate.APIAudit {
	return predicate.APIAudit(func(s *sql.Selector) {
		p(s.Not())
	})
}
