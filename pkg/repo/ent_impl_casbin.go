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

	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
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
m = g(r.sub, p.sub, r.dom) && keyMatch2(r.dom, p.dom) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
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

// CreateGrant creats a new grant in the policy database
//
// grouping policy for RBAC with domain pattern
// example: https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_policy.csv
//
// default domain admin token =>  p, hash(token1234), domain, *, *, allow
func (e *entImpl) CreateGrant(ctx context.Context, opt *CreateGrantOption) (g *ent.Grant, err error) {
	defer func() { err = e.HandleError(err) }()

	m := mergeAllowedHTTPMethods(opt.AllowedHTTPMethods)

	// If the paths is not specified, we then assume it's for all the paths "*"
	if len(opt.Paths) == 0 {
		opt.Paths = []string{"*"}
	}

	for _, path := range opt.Paths {
		_, err = e.AddPolicy(AuthzPolicy{
			Subject: opt.HashedGrantToken,
			Domain:  opt.Domain,
			Object:  path,
			Action:  m,
			Effect:  "allow",
		})

		if err != nil {
			return nil, err
		}
	}

	return &ent.Grant{
		HashedGrantToken:   opt.HashedGrantToken,
		Domain:             opt.Domain,
		Version:            opt.Version,
		AllowedHTTPMethods: m,
		Paths:              opt.Paths,
	}, nil
}

func (e *entImpl) Enforce(r AuthzRequest) (result bool, err error) {
	// sub, dom, obj, act
	result, err = e.enforcer.Enforce(
		r.Subject,
		r.Domain,
		r.Object,
		r.Action,
	)

	if err != nil {
		return false, UnauthorizedError{Err: err}
	}

	return result, nil
}

func (e *entImpl) AddPolicy(p AuthzPolicy) (bool, error) {
	// sub, dom, obj, act, eft
	added, err := e.enforcer.AddPolicy(
		p.Subject,
		p.Domain,
		p.Object,
		p.Action,
		p.Effect,
	)

	if err != nil {
		return false, ValidationError{Err: err}
	}
	return added, nil
}

func (e *entImpl) AddGroupingPolicy(gp AuthzGroupingPolicy) (bool, error) {
	// sub, group, dom
	added, err := e.enforcer.AddGroupingPolicy(
		gp.Subject,
		gp.Group,
		gp.Domain,
	)

	if err != nil {
		return false, ValidationError{Err: err}
	}
	return added, nil
}
