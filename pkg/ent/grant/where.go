// Code generated by entc, DO NOT EDIT.

package grant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-privacy/opv/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// HashedToken applies equality check predicate on the "hashed_token" field. It's identical to HashedTokenEQ.
func HashedToken(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedToken), v))
	})
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// AllowedHTTPMethods applies equality check predicate on the "allowed_http_methods" field. It's identical to AllowedHTTPMethodsEQ.
func AllowedHTTPMethods(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllowedHTTPMethods), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HashedTokenEQ applies the EQ predicate on the "hashed_token" field.
func HashedTokenEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedToken), v))
	})
}

// HashedTokenNEQ applies the NEQ predicate on the "hashed_token" field.
func HashedTokenNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashedToken), v))
	})
}

// HashedTokenIn applies the In predicate on the "hashed_token" field.
func HashedTokenIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHashedToken), v...))
	})
}

// HashedTokenNotIn applies the NotIn predicate on the "hashed_token" field.
func HashedTokenNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHashedToken), v...))
	})
}

// HashedTokenGT applies the GT predicate on the "hashed_token" field.
func HashedTokenGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashedToken), v))
	})
}

// HashedTokenGTE applies the GTE predicate on the "hashed_token" field.
func HashedTokenGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashedToken), v))
	})
}

// HashedTokenLT applies the LT predicate on the "hashed_token" field.
func HashedTokenLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashedToken), v))
	})
}

// HashedTokenLTE applies the LTE predicate on the "hashed_token" field.
func HashedTokenLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHashedToken), v))
	})
}

// HashedTokenContains applies the Contains predicate on the "hashed_token" field.
func HashedTokenContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHashedToken), v))
	})
}

// HashedTokenHasPrefix applies the HasPrefix predicate on the "hashed_token" field.
func HashedTokenHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHashedToken), v))
	})
}

// HashedTokenHasSuffix applies the HasSuffix predicate on the "hashed_token" field.
func HashedTokenHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHashedToken), v))
	})
}

// HashedTokenEqualFold applies the EqualFold predicate on the "hashed_token" field.
func HashedTokenEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHashedToken), v))
	})
}

// HashedTokenContainsFold applies the ContainsFold predicate on the "hashed_token" field.
func HashedTokenContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHashedToken), v))
	})
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	})
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func DomainNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func DomainGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	})
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	})
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	})
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	})
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	})
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	})
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	})
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	})
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVersion), v))
	})
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVersion), v))
	})
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVersion), v))
	})
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVersion), v))
	})
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVersion), v))
	})
}

// AllowedHTTPMethodsEQ applies the EQ predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsNEQ applies the NEQ predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsIn applies the In predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAllowedHTTPMethods), v...))
	})
}

// AllowedHTTPMethodsNotIn applies the NotIn predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAllowedHTTPMethods), v...))
	})
}

// AllowedHTTPMethodsGT applies the GT predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsGTE applies the GTE predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsLT applies the LT predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsLTE applies the LTE predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsContains applies the Contains predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsHasPrefix applies the HasPrefix predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsHasSuffix applies the HasSuffix predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsEqualFold applies the EqualFold predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAllowedHTTPMethods), v))
	})
}

// AllowedHTTPMethodsContainsFold applies the ContainsFold predicate on the "allowed_http_methods" field.
func AllowedHTTPMethodsContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAllowedHTTPMethods), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
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
func Not(p predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		p(s.Not())
	})
}
