package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/avast/retry-go"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
	_ "github.com/mattn/go-sqlite3"    // sqlite3 driver

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/apiaudit"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/migrate"
	"github.com/open-privacy/opv/pkg/ent/predicate"
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

func setupEntDB() (*ent.Client, *casbin.SyncedEnforcer, error) {
	var entClient *ent.Client
	var enforcer *casbin.SyncedEnforcer
	var err error

	err = retry.Do(
		func() error {
			var db *sql.DB
			switch config.ENV.DBDriver {
			case dialect.MySQL, dialect.Postgres, dialect.SQLite:
				driver, err := entsql.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
				if err != nil {
					return err
				}
				entClient = ent.NewClient(ent.Driver(driver))
				db = driver.DB()
			default:
				return fmt.Errorf("unsupported database driver %s", config.ENV.DBDriver)
			}

			// Run Ent Migration
			if err := entClient.Schema.Create(
				context.Background(),
				migrate.WithDropIndex(true),
			); err != nil {
				return fmt.Errorf("failed to migrate ent schema: %v", err)
			}

			// Run Casbin Migration
			enforcer, err = newCasbin(db)
			if err != nil {
				return fmt.Errorf("failed to create casbin enforcer: %v", err)
			}

			return nil
		},
		retry.Attempts(config.ENV.DBSetupRetryAttempts),
		retry.Delay(config.ENV.DBSetupRetryDelay),
	)
	return entClient, enforcer, err
}

func newEntImpl() (*entImpl, error) {
	entClient, enforcer, err := setupEntDB()
	if err != nil {
		return nil, err
	}
	return &entImpl{entClient: entClient, enforcer: enforcer}, nil
}

func (e *entImpl) Close() {
	e.entClient.Close()
}

func (e *entImpl) HandleError(err error) error {
	if err == nil {
		return nil
	}

	if ent.IsNotFound(err) {
		return NotFoundError{Err: err}
	}
	if ent.IsValidationError(err) {
		return ValidationError{Err: err, Message: "Validation error"}
	}
	if ent.IsConstraintError(err) {
		var errorMessage = strings.ToLower(err.Error())
		if strings.Contains(errorMessage, "unique constraint") && strings.Contains(errorMessage, "insert node to table \"facts\"") {
			return ValidationError{Err: err, Message: "fact_value already exists for this scope"}
		}
	}

	switch err.(type) {
	case govalidator.Errors, govalidator.Error, validator.ValidationErrors:
		return ValidationError{Err: err, Message: "Validation error"}
	case NotFoundError, ValidationError, UnauthorizedError:
		return err
	}

	return err
}

func (e *entImpl) Enforce(rvals ...interface{}) (result bool, err error) {
	result, err = e.enforcer.Enforce(rvals...)
	if err != nil {
		return false, UnauthorizedError{Err: err}
	}

	return result, nil
}

func (e *entImpl) AddPolicy(params ...interface{}) (bool, error) {
	added, err := e.enforcer.AddPolicy(params...)

	if err != nil {
		return false, ValidationError{Err: err}
	}
	return added, nil
}

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

// CreateGrant creats a new grant in the policy database
//
// grouping policy for RBAC with domain pattern
// example: https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_policy.csv
//
// default domain admin token =>  p, hash(token1234), domain, *, *, allow
func (e *entImpl) CreateGrant(ctx context.Context, opt *CreateGrantOption) (g *ent.Grant, err error) {
	defer func() { err = e.HandleError(err) }()

	m := mergeAllowedHTTPMethods(opt.AllowedHTTPMethods)
	_, err = e.enforcer.AddPolicy(
		opt.HashedGrantToken,
		opt.Domain,
		"*",
		m,
		"allow",
	)

	if err != nil {
		return nil, err
	}

	return &ent.Grant{
		HashedGrantToken:   opt.HashedGrantToken,
		Domain:             opt.Domain,
		Version:            opt.Version,
		AllowedHTTPMethods: m,
	}, nil
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

	return e.entClient.APIAudit.Query().Where(conds...).All(ctx)
}
