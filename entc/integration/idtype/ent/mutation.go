// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent/entc/integration/idtype/ent/predicate"
	"entgo.io/ent/entc/integration/idtype/ent/user"

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
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op               Op
	typ              string
	id               *uint64
	name             *string
	clearedFields    map[string]struct{}
	spouse           *uint64
	clearedspouse    bool
	followers        map[uint64]struct{}
	removedfollowers map[uint64]struct{}
	clearedfollowers bool
	following        map[uint64]struct{}
	removedfollowing map[uint64]struct{}
	clearedfollowing bool
	done             bool
	oldValue         func(context.Context) (*User, error)
	predicates       []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uint64) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetSpouseID sets the "spouse" edge to the User entity by id.
func (m *UserMutation) SetSpouseID(id uint64) {
	m.spouse = &id
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (m *UserMutation) ClearSpouse() {
	m.clearedspouse = true
}

// SpouseCleared reports if the "spouse" edge to the User entity was cleared.
func (m *UserMutation) SpouseCleared() bool {
	return m.clearedspouse
}

// SpouseID returns the "spouse" edge ID in the mutation.
func (m *UserMutation) SpouseID() (id uint64, exists bool) {
	if m.spouse != nil {
		return *m.spouse, true
	}
	return
}

// SpouseIDs returns the "spouse" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// SpouseID instead. It exists only for internal usage by the builders.
func (m *UserMutation) SpouseIDs() (ids []uint64) {
	if id := m.spouse; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetSpouse resets all changes to the "spouse" edge.
func (m *UserMutation) ResetSpouse() {
	m.spouse = nil
	m.clearedspouse = false
}

// AddFollowerIDs adds the "followers" edge to the User entity by ids.
func (m *UserMutation) AddFollowerIDs(ids ...uint64) {
	if m.followers == nil {
		m.followers = make(map[uint64]struct{})
	}
	for i := range ids {
		m.followers[ids[i]] = struct{}{}
	}
}

// ClearFollowers clears the "followers" edge to the User entity.
func (m *UserMutation) ClearFollowers() {
	m.clearedfollowers = true
}

// FollowersCleared reports if the "followers" edge to the User entity was cleared.
func (m *UserMutation) FollowersCleared() bool {
	return m.clearedfollowers
}

// RemoveFollowerIDs removes the "followers" edge to the User entity by IDs.
func (m *UserMutation) RemoveFollowerIDs(ids ...uint64) {
	if m.removedfollowers == nil {
		m.removedfollowers = make(map[uint64]struct{})
	}
	for i := range ids {
		delete(m.followers, ids[i])
		m.removedfollowers[ids[i]] = struct{}{}
	}
}

// RemovedFollowers returns the removed IDs of the "followers" edge to the User entity.
func (m *UserMutation) RemovedFollowersIDs() (ids []uint64) {
	for id := range m.removedfollowers {
		ids = append(ids, id)
	}
	return
}

// FollowersIDs returns the "followers" edge IDs in the mutation.
func (m *UserMutation) FollowersIDs() (ids []uint64) {
	for id := range m.followers {
		ids = append(ids, id)
	}
	return
}

// ResetFollowers resets all changes to the "followers" edge.
func (m *UserMutation) ResetFollowers() {
	m.followers = nil
	m.clearedfollowers = false
	m.removedfollowers = nil
}

// AddFollowingIDs adds the "following" edge to the User entity by ids.
func (m *UserMutation) AddFollowingIDs(ids ...uint64) {
	if m.following == nil {
		m.following = make(map[uint64]struct{})
	}
	for i := range ids {
		m.following[ids[i]] = struct{}{}
	}
}

// ClearFollowing clears the "following" edge to the User entity.
func (m *UserMutation) ClearFollowing() {
	m.clearedfollowing = true
}

// FollowingCleared reports if the "following" edge to the User entity was cleared.
func (m *UserMutation) FollowingCleared() bool {
	return m.clearedfollowing
}

// RemoveFollowingIDs removes the "following" edge to the User entity by IDs.
func (m *UserMutation) RemoveFollowingIDs(ids ...uint64) {
	if m.removedfollowing == nil {
		m.removedfollowing = make(map[uint64]struct{})
	}
	for i := range ids {
		delete(m.following, ids[i])
		m.removedfollowing[ids[i]] = struct{}{}
	}
}

// RemovedFollowing returns the removed IDs of the "following" edge to the User entity.
func (m *UserMutation) RemovedFollowingIDs() (ids []uint64) {
	for id := range m.removedfollowing {
		ids = append(ids, id)
	}
	return
}

// FollowingIDs returns the "following" edge IDs in the mutation.
func (m *UserMutation) FollowingIDs() (ids []uint64) {
	for id := range m.following {
		ids = append(ids, id)
	}
	return
}

// ResetFollowing resets all changes to the "following" edge.
func (m *UserMutation) ResetFollowing() {
	m.following = nil
	m.clearedfollowing = false
	m.removedfollowing = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.spouse != nil {
		edges = append(edges, user.EdgeSpouse)
	}
	if m.followers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.following != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeSpouse:
		if id := m.spouse; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.followers))
		for id := range m.followers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.following))
		for id := range m.following {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedfollowers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.removedfollowing != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.removedfollowers))
		for id := range m.removedfollowers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.removedfollowing))
		for id := range m.removedfollowing {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedspouse {
		edges = append(edges, user.EdgeSpouse)
	}
	if m.clearedfollowers {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.clearedfollowing {
		edges = append(edges, user.EdgeFollowing)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeSpouse:
		return m.clearedspouse
	case user.EdgeFollowers:
		return m.clearedfollowers
	case user.EdgeFollowing:
		return m.clearedfollowing
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeSpouse:
		m.ClearSpouse()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeSpouse:
		m.ResetSpouse()
		return nil
	case user.EdgeFollowers:
		m.ResetFollowers()
		return nil
	case user.EdgeFollowing:
		m.ResetFollowing()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
