// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package session

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "sessions"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "device_sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sessions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_sessions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() schema.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func([]byte) error
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeviceField orders the results by device field.
func ByDeviceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStep(), sql.OrderByField(field, opts...))
	}
}
func newDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
	)
}
