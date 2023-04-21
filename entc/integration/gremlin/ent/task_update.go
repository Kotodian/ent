// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/ent/schema/task"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TaskUpdate) SetPriority(t task.Priority) *TaskUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(t)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePriority(t *task.Priority) *TaskUpdate {
	if t != nil {
		tu.SetPriority(*t)
	}
	return tu
}

// AddPriority adds t to the "priority" field.
func (tu *TaskUpdate) AddPriority(t task.Priority) *TaskUpdate {
	tu.mutation.AddPriority(t)
	return tu
}

// SetPriorities sets the "priorities" field.
func (tu *TaskUpdate) SetPriorities(m map[string]task.Priority) *TaskUpdate {
	tu.mutation.SetPriorities(m)
	return tu
}

// ClearPriorities clears the value of the "priorities" field.
func (tu *TaskUpdate) ClearPriorities() *TaskUpdate {
	tu.mutation.ClearPriorities()
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// ClearName clears the value of the "name" field.
func (tu *TaskUpdate) ClearName() *TaskUpdate {
	tu.mutation.ClearName()
	return tu
}

// SetOwner sets the "owner" field.
func (tu *TaskUpdate) SetOwner(s string) *TaskUpdate {
	tu.mutation.SetOwner(s)
	return tu
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOwner(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOwner(*s)
	}
	return tu
}

// ClearOwner clears the value of the "owner" field.
func (tu *TaskUpdate) ClearOwner() *TaskUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// SetOrder sets the "order" field.
func (tu *TaskUpdate) SetOrder(i int) *TaskUpdate {
	tu.mutation.ResetOrder()
	tu.mutation.SetOrder(i)
	return tu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOrder(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOrder(*i)
	}
	return tu
}

// AddOrder adds i to the "order" field.
func (tu *TaskUpdate) AddOrder(i int) *TaskUpdate {
	tu.mutation.AddOrder(i)
	return tu
}

// ClearOrder clears the value of the "order" field.
func (tu *TaskUpdate) ClearOrder() *TaskUpdate {
	tu.mutation.ClearOrder()
	return tu
}

// SetOrderOption sets the "order_option" field.
func (tu *TaskUpdate) SetOrderOption(i int) *TaskUpdate {
	tu.mutation.ResetOrderOption()
	tu.mutation.SetOrderOption(i)
	return tu
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOrderOption(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOrderOption(*i)
	}
	return tu
}

// AddOrderOption adds i to the "order_option" field.
func (tu *TaskUpdate) AddOrderOption(i int) *TaskUpdate {
	tu.mutation.AddOrderOption(i)
	return tu
}

// ClearOrderOption clears the value of the "order_option" field.
func (tu *TaskUpdate) ClearOrderOption() *TaskUpdate {
	tu.mutation.ClearOrderOption()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TaskMutation](ctx, tu.gremlinSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := tu.gremlin().Query()
	if err := tu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	tu.mutation.done = true
	return res.ReadInt()
}

func (tu *TaskUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(enttask.Label)
	for _, p := range tu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value, ok := tu.mutation.Priority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, value)
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, __.Union(__.Values(enttask.FieldPriority), __.Constant(value)).Sum())
	}
	if value, ok := tu.mutation.Priorities(); ok {
		v.Property(dsl.Single, enttask.FieldPriorities, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		v.Property(dsl.Single, enttask.FieldName, value)
	}
	if value, ok := tu.mutation.Owner(); ok {
		v.Property(dsl.Single, enttask.FieldOwner, value)
	}
	if value, ok := tu.mutation.Order(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, value)
	}
	if value, ok := tu.mutation.AddedOrder(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, __.Union(__.Values(enttask.FieldOrder), __.Constant(value)).Sum())
	}
	if value, ok := tu.mutation.OrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, value)
	}
	if value, ok := tu.mutation.AddedOrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, __.Union(__.Values(enttask.FieldOrderOption), __.Constant(value)).Sum())
	}
	var properties []any
	if tu.mutation.PrioritiesCleared() {
		properties = append(properties, enttask.FieldPriorities)
	}
	if tu.mutation.NameCleared() {
		properties = append(properties, enttask.FieldName)
	}
	if tu.mutation.OwnerCleared() {
		properties = append(properties, enttask.FieldOwner)
	}
	if tu.mutation.OrderCleared() {
		properties = append(properties, enttask.FieldOrder)
	}
	if tu.mutation.OrderOptionCleared() {
		properties = append(properties, enttask.FieldOrderOption)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetPriority sets the "priority" field.
func (tuo *TaskUpdateOne) SetPriority(t task.Priority) *TaskUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(t)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePriority(t *task.Priority) *TaskUpdateOne {
	if t != nil {
		tuo.SetPriority(*t)
	}
	return tuo
}

// AddPriority adds t to the "priority" field.
func (tuo *TaskUpdateOne) AddPriority(t task.Priority) *TaskUpdateOne {
	tuo.mutation.AddPriority(t)
	return tuo
}

// SetPriorities sets the "priorities" field.
func (tuo *TaskUpdateOne) SetPriorities(m map[string]task.Priority) *TaskUpdateOne {
	tuo.mutation.SetPriorities(m)
	return tuo
}

// ClearPriorities clears the value of the "priorities" field.
func (tuo *TaskUpdateOne) ClearPriorities() *TaskUpdateOne {
	tuo.mutation.ClearPriorities()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// ClearName clears the value of the "name" field.
func (tuo *TaskUpdateOne) ClearName() *TaskUpdateOne {
	tuo.mutation.ClearName()
	return tuo
}

// SetOwner sets the "owner" field.
func (tuo *TaskUpdateOne) SetOwner(s string) *TaskUpdateOne {
	tuo.mutation.SetOwner(s)
	return tuo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOwner(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOwner(*s)
	}
	return tuo
}

// ClearOwner clears the value of the "owner" field.
func (tuo *TaskUpdateOne) ClearOwner() *TaskUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// SetOrder sets the "order" field.
func (tuo *TaskUpdateOne) SetOrder(i int) *TaskUpdateOne {
	tuo.mutation.ResetOrder()
	tuo.mutation.SetOrder(i)
	return tuo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOrder(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOrder(*i)
	}
	return tuo
}

// AddOrder adds i to the "order" field.
func (tuo *TaskUpdateOne) AddOrder(i int) *TaskUpdateOne {
	tuo.mutation.AddOrder(i)
	return tuo
}

// ClearOrder clears the value of the "order" field.
func (tuo *TaskUpdateOne) ClearOrder() *TaskUpdateOne {
	tuo.mutation.ClearOrder()
	return tuo
}

// SetOrderOption sets the "order_option" field.
func (tuo *TaskUpdateOne) SetOrderOption(i int) *TaskUpdateOne {
	tuo.mutation.ResetOrderOption()
	tuo.mutation.SetOrderOption(i)
	return tuo
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOrderOption(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOrderOption(*i)
	}
	return tuo
}

// AddOrderOption adds i to the "order_option" field.
func (tuo *TaskUpdateOne) AddOrderOption(i int) *TaskUpdateOne {
	tuo.mutation.AddOrderOption(i)
	return tuo
}

// ClearOrderOption clears the value of the "order_option" field.
func (tuo *TaskUpdateOne) ClearOrderOption() *TaskUpdateOne {
	tuo.mutation.ClearOrderOption()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks[*Task, TaskMutation](ctx, tuo.gremlinSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) gremlinSave(ctx context.Context) (*Task, error) {
	res := &gremlin.Response{}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	query, bindings := tuo.gremlin(id).Query()
	if err := tuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	tuo.mutation.done = true
	t := &Task{config: tuo.config}
	if err := t.FromResponse(res); err != nil {
		return nil, err
	}
	return t, nil
}

func (tuo *TaskUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value, ok := tuo.mutation.Priority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, value)
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, __.Union(__.Values(enttask.FieldPriority), __.Constant(value)).Sum())
	}
	if value, ok := tuo.mutation.Priorities(); ok {
		v.Property(dsl.Single, enttask.FieldPriorities, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		v.Property(dsl.Single, enttask.FieldName, value)
	}
	if value, ok := tuo.mutation.Owner(); ok {
		v.Property(dsl.Single, enttask.FieldOwner, value)
	}
	if value, ok := tuo.mutation.Order(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, value)
	}
	if value, ok := tuo.mutation.AddedOrder(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, __.Union(__.Values(enttask.FieldOrder), __.Constant(value)).Sum())
	}
	if value, ok := tuo.mutation.OrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, value)
	}
	if value, ok := tuo.mutation.AddedOrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, __.Union(__.Values(enttask.FieldOrderOption), __.Constant(value)).Sum())
	}
	var properties []any
	if tuo.mutation.PrioritiesCleared() {
		properties = append(properties, enttask.FieldPriorities)
	}
	if tuo.mutation.NameCleared() {
		properties = append(properties, enttask.FieldName)
	}
	if tuo.mutation.OwnerCleared() {
		properties = append(properties, enttask.FieldOwner)
	}
	if tuo.mutation.OrderCleared() {
		properties = append(properties, enttask.FieldOrder)
	}
	if tuo.mutation.OrderOptionCleared() {
		properties = append(properties, enttask.FieldOrderOption)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if len(tuo.fields) > 0 {
		fields := make([]any, 0, len(tuo.fields)+1)
		fields = append(fields, true)
		for _, f := range tuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
