package repo

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
	_ "github.com/mattn/go-sqlite3"    // sqlite3 driver

	"github.com/roney492/opv/pkg/ent"
	"github.com/roney492/opv/pkg/ent/apiaudit"
	"github.com/roney492/opv/pkg/ent/fact"
	"github.com/roney492/opv/pkg/ent/facttype"
	"github.com/roney492/opv/pkg/ent/predicate"
	"github.com/roney492/opv/pkg/ent/scope"
)

func (e *entImpl) CreateFact(ctx context.Context, opt *CreateFactOption) (f *ent.Fact, err error) {
	defer func() { err = e.HandleError(err) }()

	// If scope is empty (nil or custom_id is empty) then just create the fact anyway
	if opt.Scope == nil || opt.Scope.CustomID == "" {
		return e.entClient.Fact.Create().
			SetScope(opt.Scope).
			SetFactType(opt.FactType).
			SetDomain(opt.Domain).
			SetHashedValue(opt.HashedValue).
			SetEncryptedValue(opt.EncryptedValue).
			Save(ctx)
	}

	// If scope is already set, try to dedup hashed_value first
	exists, err := e.entClient.Fact.Query().Where(
		fact.DeletedAtIsNil(),
		fact.HasScopeWith(scope.CustomID(opt.Scope.CustomID)),
		fact.HasFactTypeWith(facttype.Slug(opt.FactType.Slug)),
		fact.HashedValue(opt.HashedValue),
		fact.Domain(opt.Domain),
	).Exist(ctx)

	// already exists, we shouldn't create a new fact with the same combination of [scope, fact_type, hashed_value]
	// return validation error
	if err != nil {
		return nil, err
	}
	if exists {
		err = fmt.Errorf("hashed_value already exists")
		return nil, ValidationError{Err: err, Message: err.Error()}
	}

	return e.entClient.Fact.Create().
		SetScope(opt.Scope).
		SetFactType(opt.FactType).
		SetDomain(opt.Domain).
		SetHashedValue(opt.HashedValue).
		SetEncryptedValue(opt.EncryptedValue).
		Save(ctx)
}

func (e *entImpl) GetFact(ctx context.Context, opt *GetFactOption) (f *ent.Fact, err error) {
	defer func() { err = e.HandleError(err) }()

	return e.entClient.Fact.Query().WithScope().WithFactType().Where(
		fact.DeletedAtIsNil(),
		fact.ID(opt.FactID),
		fact.Domain(opt.Domain),
	).Only(ctx)
}

func (e *entImpl) CreateFactType(ctx context.Context, opt *CreateFactTypeOption) (ft *ent.FactType, err error) {
	defer func() { err = e.HandleError(err) }()

	ft, err = e.entClient.FactType.Query().Where(
		facttype.DeletedAtIsNil(),
		facttype.Slug(opt.FactTypeSlug),
	).Only(ctx)

	if ent.IsNotFound(err) {
		ft, err = e.entClient.FactType.Create().
			SetBuiltIn(opt.BuiltIn).
			SetSlug(opt.FactTypeSlug).
			SetValidation(opt.FactTypeValidation).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return ft, nil
	}

	return ft, err
}

func (e *entImpl) GetFactTypeBySlug(ctx context.Context, slug string) (ft *ent.FactType, err error) {
	defer func() { err = e.HandleError(err) }()

	return e.entClient.FactType.Query().Where(
		facttype.DeletedAtIsNil(),
		facttype.Slug(slug),
	).Only(ctx)
}

func (e *entImpl) CreateScope(ctx context.Context, opt *CreateScopeOption) (s *ent.Scope, err error) {
	defer func() { err = e.HandleError(err) }()

	s, err = e.entClient.Scope.Query().Where(
		scope.DeletedAtIsNil(),
		scope.CustomID(opt.ScopeCustomID),
		scope.Domain(opt.Domain),
	).Only(ctx)

	if ent.IsNotFound(err) {
		return e.entClient.Scope.Create().
			SetCustomID(opt.ScopeCustomID).
			SetDomain(opt.Domain).
			Save(ctx)
	}
	return s, err
}

func (e *entImpl) GetScope(ctx context.Context, opt *GetScopeOption) (s *ent.Scope, err error) {
	defer func() { err = e.HandleError(err) }()

	return e.entClient.Scope.Query().Where(
		scope.DeletedAtIsNil(),
		scope.CustomID(opt.ScopeCustomID),
		scope.Domain(opt.Domain),
	).Only(ctx)
}

func (e *entImpl) CreateAPIAudit(ctx context.Context, opt *CreateAPIAuditOption) (a *ent.APIAudit, err error) {
	defer func() { err = e.HandleError(err) }()

	if opt.Plane == DataplaneName || opt.Plane == ControlplaneName || opt.Plane == ProxyplaneName {
		return e.entClient.APIAudit.Create().
			SetPlane(opt.Plane).
			SetNillableDomain(opt.Domain).
			SetNillableHashedGrantToken(opt.HashedGrantToken).
			SetNillableHTTPPath(opt.HTTPPath).
			SetNillableHTTPMethod(opt.HTTPMethod).
			SetNillableSentHTTPStatus(opt.SentHTTPStatus).
			Save(ctx)
	}

	return nil, fmt.Errorf("not supported plane for audit logs, %s", opt.Plane)
}

func (e *entImpl) QueryAPIAudits(ctx context.Context, opt *QueryAPIAuditOption) (a []*ent.APIAudit, err error) {
	defer func() { err = e.HandleError(err) }()

	conds := []predicate.APIAudit{}

	if opt.Plane != nil {
		conds = append(conds, apiaudit.Plane(*opt.Plane))
	}
	if opt.Domain != nil {
		conds = append(conds, apiaudit.Domain(*opt.Domain))
	}
	if opt.HashedGrantToken != nil {
		conds = append(conds, apiaudit.HashedGrantToken(*opt.HashedGrantToken))
	}
	if opt.HTTPMethod != nil {
		conds = append(conds, apiaudit.HTTPMethod(*opt.HTTPMethod))
	}
	if opt.HTTPPath != nil {
		conds = append(conds, apiaudit.HTTPPath(*opt.HTTPPath))
	}
	if opt.SentHTTPStatus != nil {
		conds = append(conds, apiaudit.SentHTTPStatus(*opt.SentHTTPStatus))
	}

	query := e.entClient.APIAudit.Query().Where(conds...)

	if opt.Limit != nil {
		query = query.Limit(*opt.Limit)
	}
	if opt.Offset != nil {
		query = query.Offset(*opt.Offset)
	}
	if opt.OrderBy != nil {
		if opt.OrderDesc {
			query = query.Order(ent.Desc(*opt.OrderBy))
		} else {
			query = query.Order(ent.Asc(*opt.OrderBy))
		}
	}

	return query.All(ctx)
}
