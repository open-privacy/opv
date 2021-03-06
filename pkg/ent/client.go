// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/open-privacy/opv/pkg/ent/migrate"

	"github.com/open-privacy/opv/pkg/ent/apiaudit"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/grant"
	"github.com/open-privacy/opv/pkg/ent/scope"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// APIAudit is the client for interacting with the APIAudit builders.
	APIAudit *APIAuditClient
	// Fact is the client for interacting with the Fact builders.
	Fact *FactClient
	// FactType is the client for interacting with the FactType builders.
	FactType *FactTypeClient
	// Grant is the client for interacting with the Grant builders.
	Grant *GrantClient
	// Scope is the client for interacting with the Scope builders.
	Scope *ScopeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.APIAudit = NewAPIAuditClient(c.config)
	c.Fact = NewFactClient(c.config)
	c.FactType = NewFactTypeClient(c.config)
	c.Grant = NewGrantClient(c.config)
	c.Scope = NewScopeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		APIAudit: NewAPIAuditClient(cfg),
		Fact:     NewFactClient(cfg),
		FactType: NewFactTypeClient(cfg),
		Grant:    NewGrantClient(cfg),
		Scope:    NewScopeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		APIAudit: NewAPIAuditClient(cfg),
		Fact:     NewFactClient(cfg),
		FactType: NewFactTypeClient(cfg),
		Grant:    NewGrantClient(cfg),
		Scope:    NewScopeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		APIAudit.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.APIAudit.Use(hooks...)
	c.Fact.Use(hooks...)
	c.FactType.Use(hooks...)
	c.Grant.Use(hooks...)
	c.Scope.Use(hooks...)
}

// APIAuditClient is a client for the APIAudit schema.
type APIAuditClient struct {
	config
}

// NewAPIAuditClient returns a client for the APIAudit from the given config.
func NewAPIAuditClient(c config) *APIAuditClient {
	return &APIAuditClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apiaudit.Hooks(f(g(h())))`.
func (c *APIAuditClient) Use(hooks ...Hook) {
	c.hooks.APIAudit = append(c.hooks.APIAudit, hooks...)
}

// Create returns a create builder for APIAudit.
func (c *APIAuditClient) Create() *APIAuditCreate {
	mutation := newAPIAuditMutation(c.config, OpCreate)
	return &APIAuditCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of APIAudit entities.
func (c *APIAuditClient) CreateBulk(builders ...*APIAuditCreate) *APIAuditCreateBulk {
	return &APIAuditCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for APIAudit.
func (c *APIAuditClient) Update() *APIAuditUpdate {
	mutation := newAPIAuditMutation(c.config, OpUpdate)
	return &APIAuditUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *APIAuditClient) UpdateOne(aa *APIAudit) *APIAuditUpdateOne {
	mutation := newAPIAuditMutation(c.config, OpUpdateOne, withAPIAudit(aa))
	return &APIAuditUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *APIAuditClient) UpdateOneID(id string) *APIAuditUpdateOne {
	mutation := newAPIAuditMutation(c.config, OpUpdateOne, withAPIAuditID(id))
	return &APIAuditUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for APIAudit.
func (c *APIAuditClient) Delete() *APIAuditDelete {
	mutation := newAPIAuditMutation(c.config, OpDelete)
	return &APIAuditDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *APIAuditClient) DeleteOne(aa *APIAudit) *APIAuditDeleteOne {
	return c.DeleteOneID(aa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *APIAuditClient) DeleteOneID(id string) *APIAuditDeleteOne {
	builder := c.Delete().Where(apiaudit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &APIAuditDeleteOne{builder}
}

// Query returns a query builder for APIAudit.
func (c *APIAuditClient) Query() *APIAuditQuery {
	return &APIAuditQuery{config: c.config}
}

// Get returns a APIAudit entity by its id.
func (c *APIAuditClient) Get(ctx context.Context, id string) (*APIAudit, error) {
	return c.Query().Where(apiaudit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *APIAuditClient) GetX(ctx context.Context, id string) *APIAudit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *APIAuditClient) Hooks() []Hook {
	return c.hooks.APIAudit
}

// FactClient is a client for the Fact schema.
type FactClient struct {
	config
}

// NewFactClient returns a client for the Fact from the given config.
func NewFactClient(c config) *FactClient {
	return &FactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fact.Hooks(f(g(h())))`.
func (c *FactClient) Use(hooks ...Hook) {
	c.hooks.Fact = append(c.hooks.Fact, hooks...)
}

// Create returns a create builder for Fact.
func (c *FactClient) Create() *FactCreate {
	mutation := newFactMutation(c.config, OpCreate)
	return &FactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Fact entities.
func (c *FactClient) CreateBulk(builders ...*FactCreate) *FactCreateBulk {
	return &FactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Fact.
func (c *FactClient) Update() *FactUpdate {
	mutation := newFactMutation(c.config, OpUpdate)
	return &FactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FactClient) UpdateOne(f *Fact) *FactUpdateOne {
	mutation := newFactMutation(c.config, OpUpdateOne, withFact(f))
	return &FactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FactClient) UpdateOneID(id string) *FactUpdateOne {
	mutation := newFactMutation(c.config, OpUpdateOne, withFactID(id))
	return &FactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Fact.
func (c *FactClient) Delete() *FactDelete {
	mutation := newFactMutation(c.config, OpDelete)
	return &FactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FactClient) DeleteOne(f *Fact) *FactDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FactClient) DeleteOneID(id string) *FactDeleteOne {
	builder := c.Delete().Where(fact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FactDeleteOne{builder}
}

// Query returns a query builder for Fact.
func (c *FactClient) Query() *FactQuery {
	return &FactQuery{config: c.config}
}

// Get returns a Fact entity by its id.
func (c *FactClient) Get(ctx context.Context, id string) (*Fact, error) {
	return c.Query().Where(fact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FactClient) GetX(ctx context.Context, id string) *Fact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryScope queries the scope edge of a Fact.
func (c *FactClient) QueryScope(f *Fact) *ScopeQuery {
	query := &ScopeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(fact.Table, fact.FieldID, id),
			sqlgraph.To(scope.Table, scope.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, fact.ScopeTable, fact.ScopeColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFactType queries the fact_type edge of a Fact.
func (c *FactClient) QueryFactType(f *Fact) *FactTypeQuery {
	query := &FactTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(fact.Table, fact.FieldID, id),
			sqlgraph.To(facttype.Table, facttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, fact.FactTypeTable, fact.FactTypeColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FactClient) Hooks() []Hook {
	return c.hooks.Fact
}

// FactTypeClient is a client for the FactType schema.
type FactTypeClient struct {
	config
}

// NewFactTypeClient returns a client for the FactType from the given config.
func NewFactTypeClient(c config) *FactTypeClient {
	return &FactTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `facttype.Hooks(f(g(h())))`.
func (c *FactTypeClient) Use(hooks ...Hook) {
	c.hooks.FactType = append(c.hooks.FactType, hooks...)
}

// Create returns a create builder for FactType.
func (c *FactTypeClient) Create() *FactTypeCreate {
	mutation := newFactTypeMutation(c.config, OpCreate)
	return &FactTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FactType entities.
func (c *FactTypeClient) CreateBulk(builders ...*FactTypeCreate) *FactTypeCreateBulk {
	return &FactTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FactType.
func (c *FactTypeClient) Update() *FactTypeUpdate {
	mutation := newFactTypeMutation(c.config, OpUpdate)
	return &FactTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FactTypeClient) UpdateOne(ft *FactType) *FactTypeUpdateOne {
	mutation := newFactTypeMutation(c.config, OpUpdateOne, withFactType(ft))
	return &FactTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FactTypeClient) UpdateOneID(id string) *FactTypeUpdateOne {
	mutation := newFactTypeMutation(c.config, OpUpdateOne, withFactTypeID(id))
	return &FactTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FactType.
func (c *FactTypeClient) Delete() *FactTypeDelete {
	mutation := newFactTypeMutation(c.config, OpDelete)
	return &FactTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FactTypeClient) DeleteOne(ft *FactType) *FactTypeDeleteOne {
	return c.DeleteOneID(ft.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FactTypeClient) DeleteOneID(id string) *FactTypeDeleteOne {
	builder := c.Delete().Where(facttype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FactTypeDeleteOne{builder}
}

// Query returns a query builder for FactType.
func (c *FactTypeClient) Query() *FactTypeQuery {
	return &FactTypeQuery{config: c.config}
}

// Get returns a FactType entity by its id.
func (c *FactTypeClient) Get(ctx context.Context, id string) (*FactType, error) {
	return c.Query().Where(facttype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FactTypeClient) GetX(ctx context.Context, id string) *FactType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFacts queries the facts edge of a FactType.
func (c *FactTypeClient) QueryFacts(ft *FactType) *FactQuery {
	query := &FactQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ft.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(facttype.Table, facttype.FieldID, id),
			sqlgraph.To(fact.Table, fact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, facttype.FactsTable, facttype.FactsColumn),
		)
		fromV = sqlgraph.Neighbors(ft.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FactTypeClient) Hooks() []Hook {
	return c.hooks.FactType
}

// GrantClient is a client for the Grant schema.
type GrantClient struct {
	config
}

// NewGrantClient returns a client for the Grant from the given config.
func NewGrantClient(c config) *GrantClient {
	return &GrantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `grant.Hooks(f(g(h())))`.
func (c *GrantClient) Use(hooks ...Hook) {
	c.hooks.Grant = append(c.hooks.Grant, hooks...)
}

// Create returns a create builder for Grant.
func (c *GrantClient) Create() *GrantCreate {
	mutation := newGrantMutation(c.config, OpCreate)
	return &GrantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Grant entities.
func (c *GrantClient) CreateBulk(builders ...*GrantCreate) *GrantCreateBulk {
	return &GrantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Grant.
func (c *GrantClient) Update() *GrantUpdate {
	mutation := newGrantMutation(c.config, OpUpdate)
	return &GrantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GrantClient) UpdateOne(gr *Grant) *GrantUpdateOne {
	mutation := newGrantMutation(c.config, OpUpdateOne, withGrant(gr))
	return &GrantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GrantClient) UpdateOneID(id string) *GrantUpdateOne {
	mutation := newGrantMutation(c.config, OpUpdateOne, withGrantID(id))
	return &GrantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Grant.
func (c *GrantClient) Delete() *GrantDelete {
	mutation := newGrantMutation(c.config, OpDelete)
	return &GrantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GrantClient) DeleteOne(gr *Grant) *GrantDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GrantClient) DeleteOneID(id string) *GrantDeleteOne {
	builder := c.Delete().Where(grant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GrantDeleteOne{builder}
}

// Query returns a query builder for Grant.
func (c *GrantClient) Query() *GrantQuery {
	return &GrantQuery{config: c.config}
}

// Get returns a Grant entity by its id.
func (c *GrantClient) Get(ctx context.Context, id string) (*Grant, error) {
	return c.Query().Where(grant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GrantClient) GetX(ctx context.Context, id string) *Grant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GrantClient) Hooks() []Hook {
	return c.hooks.Grant
}

// ScopeClient is a client for the Scope schema.
type ScopeClient struct {
	config
}

// NewScopeClient returns a client for the Scope from the given config.
func NewScopeClient(c config) *ScopeClient {
	return &ScopeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `scope.Hooks(f(g(h())))`.
func (c *ScopeClient) Use(hooks ...Hook) {
	c.hooks.Scope = append(c.hooks.Scope, hooks...)
}

// Create returns a create builder for Scope.
func (c *ScopeClient) Create() *ScopeCreate {
	mutation := newScopeMutation(c.config, OpCreate)
	return &ScopeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Scope entities.
func (c *ScopeClient) CreateBulk(builders ...*ScopeCreate) *ScopeCreateBulk {
	return &ScopeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Scope.
func (c *ScopeClient) Update() *ScopeUpdate {
	mutation := newScopeMutation(c.config, OpUpdate)
	return &ScopeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScopeClient) UpdateOne(s *Scope) *ScopeUpdateOne {
	mutation := newScopeMutation(c.config, OpUpdateOne, withScope(s))
	return &ScopeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScopeClient) UpdateOneID(id string) *ScopeUpdateOne {
	mutation := newScopeMutation(c.config, OpUpdateOne, withScopeID(id))
	return &ScopeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Scope.
func (c *ScopeClient) Delete() *ScopeDelete {
	mutation := newScopeMutation(c.config, OpDelete)
	return &ScopeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ScopeClient) DeleteOne(s *Scope) *ScopeDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ScopeClient) DeleteOneID(id string) *ScopeDeleteOne {
	builder := c.Delete().Where(scope.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScopeDeleteOne{builder}
}

// Query returns a query builder for Scope.
func (c *ScopeClient) Query() *ScopeQuery {
	return &ScopeQuery{config: c.config}
}

// Get returns a Scope entity by its id.
func (c *ScopeClient) Get(ctx context.Context, id string) (*Scope, error) {
	return c.Query().Where(scope.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScopeClient) GetX(ctx context.Context, id string) *Scope {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFacts queries the facts edge of a Scope.
func (c *ScopeClient) QueryFacts(s *Scope) *FactQuery {
	query := &FactQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(scope.Table, scope.FieldID, id),
			sqlgraph.To(fact.Table, fact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scope.FactsTable, scope.FactsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScopeClient) Hooks() []Hook {
	return c.hooks.Scope
}
