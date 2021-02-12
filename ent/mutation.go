// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/open-privacy-vault/opv/ent/fact"
	"github.com/open-privacy-vault/opv/ent/predicate"
	"github.com/open-privacy-vault/opv/ent/scope"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFact  = "Fact"
	TypeScope = "Scope"
)

// FactMutation represents an operation that mutates the Fact nodes in the graph.
type FactMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	created_at      *time.Time
	updated_at      *time.Time
	encrypted_value *[]byte
	clearedFields   map[string]struct{}
	scope           *uuid.UUID
	clearedscope    bool
	done            bool
	oldValue        func(context.Context) (*Fact, error)
	predicates      []predicate.Fact
}

var _ ent.Mutation = (*FactMutation)(nil)

// factOption allows management of the mutation configuration using functional options.
type factOption func(*FactMutation)

// newFactMutation creates new mutation for the Fact entity.
func newFactMutation(c config, op Op, opts ...factOption) *FactMutation {
	m := &FactMutation{
		config:        c,
		op:            op,
		typ:           TypeFact,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFactID sets the ID field of the mutation.
func withFactID(id uuid.UUID) factOption {
	return func(m *FactMutation) {
		var (
			err   error
			once  sync.Once
			value *Fact
		)
		m.oldValue = func(ctx context.Context) (*Fact, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Fact.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFact sets the old Fact of the mutation.
func withFact(node *Fact) factOption {
	return func(m *FactMutation) {
		m.oldValue = func(context.Context) (*Fact, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FactMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FactMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Fact entities.
func (m *FactMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *FactMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreatedAt sets the "created_at" field.
func (m *FactMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *FactMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Fact entity.
// If the Fact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FactMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *FactMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *FactMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *FactMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Fact entity.
// If the Fact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FactMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *FactMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetEncryptedValue sets the "encrypted_value" field.
func (m *FactMutation) SetEncryptedValue(b []byte) {
	m.encrypted_value = &b
}

// EncryptedValue returns the value of the "encrypted_value" field in the mutation.
func (m *FactMutation) EncryptedValue() (r []byte, exists bool) {
	v := m.encrypted_value
	if v == nil {
		return
	}
	return *v, true
}

// OldEncryptedValue returns the old "encrypted_value" field's value of the Fact entity.
// If the Fact object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FactMutation) OldEncryptedValue(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEncryptedValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEncryptedValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEncryptedValue: %w", err)
	}
	return oldValue.EncryptedValue, nil
}

// ResetEncryptedValue resets all changes to the "encrypted_value" field.
func (m *FactMutation) ResetEncryptedValue() {
	m.encrypted_value = nil
}

// SetScopeID sets the "scope" edge to the Scope entity by id.
func (m *FactMutation) SetScopeID(id uuid.UUID) {
	m.scope = &id
}

// ClearScope clears the "scope" edge to the Scope entity.
func (m *FactMutation) ClearScope() {
	m.clearedscope = true
}

// ScopeCleared returns if the "scope" edge to the Scope entity was cleared.
func (m *FactMutation) ScopeCleared() bool {
	return m.clearedscope
}

// ScopeID returns the "scope" edge ID in the mutation.
func (m *FactMutation) ScopeID() (id uuid.UUID, exists bool) {
	if m.scope != nil {
		return *m.scope, true
	}
	return
}

// ScopeIDs returns the "scope" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ScopeID instead. It exists only for internal usage by the builders.
func (m *FactMutation) ScopeIDs() (ids []uuid.UUID) {
	if id := m.scope; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetScope resets all changes to the "scope" edge.
func (m *FactMutation) ResetScope() {
	m.scope = nil
	m.clearedscope = false
}

// Op returns the operation name.
func (m *FactMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Fact).
func (m *FactMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FactMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.created_at != nil {
		fields = append(fields, fact.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, fact.FieldUpdatedAt)
	}
	if m.encrypted_value != nil {
		fields = append(fields, fact.FieldEncryptedValue)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FactMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case fact.FieldCreatedAt:
		return m.CreatedAt()
	case fact.FieldUpdatedAt:
		return m.UpdatedAt()
	case fact.FieldEncryptedValue:
		return m.EncryptedValue()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FactMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case fact.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case fact.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case fact.FieldEncryptedValue:
		return m.OldEncryptedValue(ctx)
	}
	return nil, fmt.Errorf("unknown Fact field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FactMutation) SetField(name string, value ent.Value) error {
	switch name {
	case fact.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case fact.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case fact.FieldEncryptedValue:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEncryptedValue(v)
		return nil
	}
	return fmt.Errorf("unknown Fact field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FactMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FactMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FactMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Fact numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FactMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FactMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FactMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Fact nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FactMutation) ResetField(name string) error {
	switch name {
	case fact.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case fact.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case fact.FieldEncryptedValue:
		m.ResetEncryptedValue()
		return nil
	}
	return fmt.Errorf("unknown Fact field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FactMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.scope != nil {
		edges = append(edges, fact.EdgeScope)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FactMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case fact.EdgeScope:
		if id := m.scope; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FactMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FactMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FactMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedscope {
		edges = append(edges, fact.EdgeScope)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FactMutation) EdgeCleared(name string) bool {
	switch name {
	case fact.EdgeScope:
		return m.clearedscope
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FactMutation) ClearEdge(name string) error {
	switch name {
	case fact.EdgeScope:
		m.ClearScope()
		return nil
	}
	return fmt.Errorf("unknown Fact unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FactMutation) ResetEdge(name string) error {
	switch name {
	case fact.EdgeScope:
		m.ResetScope()
		return nil
	}
	return fmt.Errorf("unknown Fact edge %s", name)
}

// ScopeMutation represents an operation that mutates the Scope nodes in the graph.
type ScopeMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	nonce         *uuid.UUID
	expires_at    *time.Time
	clearedFields map[string]struct{}
	facts         map[uuid.UUID]struct{}
	removedfacts  map[uuid.UUID]struct{}
	clearedfacts  bool
	done          bool
	oldValue      func(context.Context) (*Scope, error)
	predicates    []predicate.Scope
}

var _ ent.Mutation = (*ScopeMutation)(nil)

// scopeOption allows management of the mutation configuration using functional options.
type scopeOption func(*ScopeMutation)

// newScopeMutation creates new mutation for the Scope entity.
func newScopeMutation(c config, op Op, opts ...scopeOption) *ScopeMutation {
	m := &ScopeMutation{
		config:        c,
		op:            op,
		typ:           TypeScope,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withScopeID sets the ID field of the mutation.
func withScopeID(id uuid.UUID) scopeOption {
	return func(m *ScopeMutation) {
		var (
			err   error
			once  sync.Once
			value *Scope
		)
		m.oldValue = func(ctx context.Context) (*Scope, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Scope.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withScope sets the old Scope of the mutation.
func withScope(node *Scope) scopeOption {
	return func(m *ScopeMutation) {
		m.oldValue = func(context.Context) (*Scope, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ScopeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ScopeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Scope entities.
func (m *ScopeMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ScopeMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreatedAt sets the "created_at" field.
func (m *ScopeMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ScopeMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Scope entity.
// If the Scope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ScopeMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ScopeMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ScopeMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ScopeMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Scope entity.
// If the Scope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ScopeMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ScopeMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetNonce sets the "nonce" field.
func (m *ScopeMutation) SetNonce(u uuid.UUID) {
	m.nonce = &u
}

// Nonce returns the value of the "nonce" field in the mutation.
func (m *ScopeMutation) Nonce() (r uuid.UUID, exists bool) {
	v := m.nonce
	if v == nil {
		return
	}
	return *v, true
}

// OldNonce returns the old "nonce" field's value of the Scope entity.
// If the Scope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ScopeMutation) OldNonce(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNonce is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNonce requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNonce: %w", err)
	}
	return oldValue.Nonce, nil
}

// ResetNonce resets all changes to the "nonce" field.
func (m *ScopeMutation) ResetNonce() {
	m.nonce = nil
}

// SetExpiresAt sets the "expires_at" field.
func (m *ScopeMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the value of the "expires_at" field in the mutation.
func (m *ScopeMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old "expires_at" field's value of the Scope entity.
// If the Scope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ScopeMutation) OldExpiresAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (m *ScopeMutation) ClearExpiresAt() {
	m.expires_at = nil
	m.clearedFields[scope.FieldExpiresAt] = struct{}{}
}

// ExpiresAtCleared returns if the "expires_at" field was cleared in this mutation.
func (m *ScopeMutation) ExpiresAtCleared() bool {
	_, ok := m.clearedFields[scope.FieldExpiresAt]
	return ok
}

// ResetExpiresAt resets all changes to the "expires_at" field.
func (m *ScopeMutation) ResetExpiresAt() {
	m.expires_at = nil
	delete(m.clearedFields, scope.FieldExpiresAt)
}

// AddFactIDs adds the "facts" edge to the Fact entity by ids.
func (m *ScopeMutation) AddFactIDs(ids ...uuid.UUID) {
	if m.facts == nil {
		m.facts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.facts[ids[i]] = struct{}{}
	}
}

// ClearFacts clears the "facts" edge to the Fact entity.
func (m *ScopeMutation) ClearFacts() {
	m.clearedfacts = true
}

// FactsCleared returns if the "facts" edge to the Fact entity was cleared.
func (m *ScopeMutation) FactsCleared() bool {
	return m.clearedfacts
}

// RemoveFactIDs removes the "facts" edge to the Fact entity by IDs.
func (m *ScopeMutation) RemoveFactIDs(ids ...uuid.UUID) {
	if m.removedfacts == nil {
		m.removedfacts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedfacts[ids[i]] = struct{}{}
	}
}

// RemovedFacts returns the removed IDs of the "facts" edge to the Fact entity.
func (m *ScopeMutation) RemovedFactsIDs() (ids []uuid.UUID) {
	for id := range m.removedfacts {
		ids = append(ids, id)
	}
	return
}

// FactsIDs returns the "facts" edge IDs in the mutation.
func (m *ScopeMutation) FactsIDs() (ids []uuid.UUID) {
	for id := range m.facts {
		ids = append(ids, id)
	}
	return
}

// ResetFacts resets all changes to the "facts" edge.
func (m *ScopeMutation) ResetFacts() {
	m.facts = nil
	m.clearedfacts = false
	m.removedfacts = nil
}

// Op returns the operation name.
func (m *ScopeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Scope).
func (m *ScopeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ScopeMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, scope.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, scope.FieldUpdatedAt)
	}
	if m.nonce != nil {
		fields = append(fields, scope.FieldNonce)
	}
	if m.expires_at != nil {
		fields = append(fields, scope.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ScopeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case scope.FieldCreatedAt:
		return m.CreatedAt()
	case scope.FieldUpdatedAt:
		return m.UpdatedAt()
	case scope.FieldNonce:
		return m.Nonce()
	case scope.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ScopeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case scope.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case scope.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case scope.FieldNonce:
		return m.OldNonce(ctx)
	case scope.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown Scope field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ScopeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case scope.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case scope.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case scope.FieldNonce:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNonce(v)
		return nil
	case scope.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown Scope field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ScopeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ScopeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ScopeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Scope numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ScopeMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(scope.FieldExpiresAt) {
		fields = append(fields, scope.FieldExpiresAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ScopeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ScopeMutation) ClearField(name string) error {
	switch name {
	case scope.FieldExpiresAt:
		m.ClearExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Scope nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ScopeMutation) ResetField(name string) error {
	switch name {
	case scope.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case scope.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case scope.FieldNonce:
		m.ResetNonce()
		return nil
	case scope.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Scope field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ScopeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.facts != nil {
		edges = append(edges, scope.EdgeFacts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ScopeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case scope.EdgeFacts:
		ids := make([]ent.Value, 0, len(m.facts))
		for id := range m.facts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ScopeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedfacts != nil {
		edges = append(edges, scope.EdgeFacts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ScopeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case scope.EdgeFacts:
		ids := make([]ent.Value, 0, len(m.removedfacts))
		for id := range m.removedfacts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ScopeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedfacts {
		edges = append(edges, scope.EdgeFacts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ScopeMutation) EdgeCleared(name string) bool {
	switch name {
	case scope.EdgeFacts:
		return m.clearedfacts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ScopeMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Scope unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ScopeMutation) ResetEdge(name string) error {
	switch name {
	case scope.EdgeFacts:
		m.ResetFacts()
		return nil
	}
	return fmt.Errorf("unknown Scope edge %s", name)
}
