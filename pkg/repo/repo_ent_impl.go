package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
	_ "github.com/mattn/go-sqlite3"    // sqlite3 driver

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/migrate"
	"github.com/open-privacy/opv/pkg/ent/scope"
)

const defaultCasbinModel = `
# RBAC with domain pattern model
# https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_model.conf

[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act, eft

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub, r.dom) && keyMatch(r.dom, p.dom) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`

func newCasbin(db *sql.DB) (*casbin.SyncedEnforcer, error) {
	a, err := sqladapter.NewAdapter(db, config.ENV.DBDriver, "casbin_rule")
	if err != nil {
		return nil, err
	}

	m, err := model.NewModelFromString(defaultCasbinModel)
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	e.StartAutoLoadPolicy(config.ENV.AuthzCasbinAutoloadInterval)
	return e, nil
}

func mergeAllowedHTTPMethods(methods []string) string {
	regexActions := []string{}
	for _, action := range methods {
		if action == "*" {
			return ".*"
		}
		regexActions = append(regexActions, fmt.Sprintf("(%s)", action))
	}
	return strings.Join(regexActions, "|")
}

type entImpl struct {
	entClient *ent.Client
	enforcer  *casbin.SyncedEnforcer
}

func newEntImpl() (*entImpl, error) {
	var entClient *ent.Client
	var db *sql.DB

	// Pick DB driver
	switch config.ENV.DBDriver {
	case dialect.MySQL, dialect.Postgres:
		driver, err := entsql.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
		if err != nil {
			return nil, fmt.Errorf("failed to open database connection: %v", err)
		}
		entClient = ent.NewClient(ent.Driver(driver))
		db = driver.DB()
	case dialect.SQLite:
		driver, err := entsql.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
		if err != nil {
			return nil, fmt.Errorf("failed to open database connection: %v", err)
		}
		entClient = ent.NewClient(ent.Driver(driver))
		db = driver.DB()
		db.SetMaxOpenConns(1)
	default:
		return nil, fmt.Errorf("unsupported database driver %s", config.ENV.DBDriver)
	}

	// Run Ent Migration
	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
	); err != nil {
		return nil, fmt.Errorf("failed to migrate ent schema: %v", err)
	}

	// Run Casbin Migration
	enforcer, err := newCasbin(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin enforcer: %v", err)
	}

	return &entImpl{entClient: entClient, enforcer: enforcer}, nil
}

func (e *entImpl) Close() {
	e.entClient.Close()
}

func (e *entImpl) Enforce(rvals ...interface{}) (bool, error) {
	return e.enforcer.Enforce(rvals...)
}

func (e *entImpl) AddPolicy(params ...interface{}) (bool, error) {
	return e.enforcer.AddPolicy(params...)
}

func (e *entImpl) CreateFact(ctx context.Context, opt *CreateFactOption) (*ent.Fact, error) {
	return e.entClient.Fact.Create().
		SetScope(opt.Scope).
		SetFactType(opt.FactType).
		SetDomain(opt.Domain).
		SetHashedValue(opt.HashedValue).
		SetEncryptedValue(opt.EncryptedValue).
		Save(ctx)
}

func (e *entImpl) GetFact(ctx context.Context, opt *GetFactOption) (*ent.Fact, error) {
	return e.entClient.Fact.Query().WithScope().WithFactType().Where(
		fact.ID(opt.FactID),
		fact.Domain(opt.Domain),
	).Only(ctx)
}

func (e *entImpl) CreateFactType(ctx context.Context, opt *CreateFactTypeOption) (*ent.FactType, error) {
	ft, err := e.GetFactTypeBySlug(ctx, opt.FactTypeSlug)
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

func (e *entImpl) GetFactTypeBySlug(ctx context.Context, slug string) (*ent.FactType, error) {
	return e.entClient.FactType.Query().Where(facttype.Slug(slug)).Only(ctx)
}

func (e *entImpl) CreateScope(ctx context.Context, opt *CreateScopeOption) (*ent.Scope, error) {
	s, err := e.GetScope(ctx, &GetScopeOption{
		ScopeCustomID: opt.ScopeCustomID,
		Domain:        opt.Domain,
	})

	if ent.IsNotFound(err) {
		s, err = e.entClient.Scope.Create().
			SetCustomID(opt.ScopeCustomID).
			SetDomain(opt.Domain).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	return s, err
}

func (e *entImpl) GetScope(ctx context.Context, opt *GetScopeOption) (*ent.Scope, error) {
	return e.entClient.Scope.Query().Where(
		scope.CustomID(opt.ScopeCustomID),
		scope.Domain(opt.Domain),
	).Only(ctx)
}

func (e *entImpl) CreateGrant(ctx context.Context, opt *CreateGrantOption) (*ent.Grant, error) {
	// grouping policy for RBAC with domain pattern
	// example: https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_policy.csv
	//
	// default domain admin token =>  p, hash(token1234), domain, *, *, allow

	m := mergeAllowedHTTPMethods(opt.AllowedHTTPMethods)
	_, err := e.enforcer.AddPolicy(
		opt.HashedToken,
		opt.Domain,
		"*",
		m,
		"allow",
	)

	if err != nil {
		return nil, err
	}

	return &ent.Grant{
		HashedToken:        opt.HashedToken,
		Domain:             opt.Domain,
		Version:            opt.Version,
		AllowedHTTPMethods: m,
	}, nil
}
