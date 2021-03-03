package repo

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
)

// Enforcer is an interface that enforces the authz access
// e.g. Enforce(sub, dom, obj, act)
type Enforcer interface {
	AddPolicy(params ...interface{}) (bool, error)
	Enforce(rvals ...interface{}) (bool, error)
}

// Repo is a set of repositories
type Repo interface {
	FactRepo
	FactTypeRepo
	ScopeRepo
	GrantRepo

	Close()
}

// CreateFactOption ...
type CreateFactOption struct {
	Domain         string
	HashedValue    string
	EncryptedValue string
	Scope          *ent.Scope
	FactType       *ent.FactType
}

// GetFactOption ...
type GetFactOption struct {
	FactID string
	Domain string
}

// CreateFactTypeOption ...
type CreateFactTypeOption struct {
	FactTypeSlug       string
	FactTypeValidation string
	BuiltIn            bool
}

// CreateScopeOption ...
type CreateScopeOption struct {
	ScopeCustomID string
	Domain        string
}

// GetScopeOption ...
type GetScopeOption struct {
	ScopeCustomID string
	Domain        string
}

// CreateGrantOption ...
type CreateGrantOption struct {
	HashedToken        string
	Domain             string
	Version            string
	AllowedHTTPMethods []string
}

// FactRepo is a repository for Fact
type FactRepo interface {
	CreateFact(ctx context.Context, opt *CreateFactOption) (*ent.Fact, error)
	GetFact(ctx context.Context, opt *GetFactOption) (*ent.Fact, error)
}

// FactTypeRepo is a repository for FactType
type FactTypeRepo interface {
	CreateFactType(ctx context.Context, opt *CreateFactTypeOption) (*ent.FactType, error)
	GetFactTypeBySlug(ctx context.Context, slug string) (*ent.FactType, error)
}

// ScopeRepo is a repository for Scope
type ScopeRepo interface {
	CreateScope(ctx context.Context, opt *CreateScopeOption) (*ent.Scope, error)
	GetScope(ctx context.Context, opt *GetScopeOption) (*ent.Scope, error)
}

// GrantRepo is a repository for Grant
type GrantRepo interface {
	CreateGrant(ctx context.Context, opt *CreateGrantOption) (*ent.Grant, error)
}

// NewRepoEnforcer creates a new RepoEnforcer
func NewRepoEnforcer() (Repo, Enforcer, error) {
	switch config.ENV.DBDriver {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		re, err := newEntImpl()
		if err != nil {
			return nil, nil, err
		}
		return re, re, nil
	}
	return nil, nil, fmt.Errorf("unsupported database driver %s", config.ENV.DBDriver)
}
