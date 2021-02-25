package authz

import (
	"database/sql"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/open-privacy/opv/pkg/config"
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

// MustNewCasbin creates the casbin enforcer with the connected sql.DB
func MustNewCasbin(db *sql.DB) *casbin.SyncedEnforcer {
	e, err := NewCasbin(db)
	if err != nil {
		panic(err)
	}
	return e
}

// NewCasbin creates the casbin enforcer with the connected sql.DB
func NewCasbin(db *sql.DB) (*casbin.SyncedEnforcer, error) {
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
