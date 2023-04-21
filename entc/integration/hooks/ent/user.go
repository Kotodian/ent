// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/hooks/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Worth holds the value of the "worth" field.
	Worth uint `json:"worth,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges            UserEdges `json:"edges"`
	user_best_friend *int
	selectValues     sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// Pets holds the value of the pets edge.
	Pets []*Pet `json:"pets,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// BestFriend holds the value of the best_friend edge.
	BestFriend *User `json:"best_friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsOrErr() ([]*Pet, error) {
	if e.loadedTypes[1] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BestFriendOrErr returns the BestFriend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) BestFriendOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.BestFriend == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BestFriend, nil
	}
	return nil, &NotLoadedError{edge: "best_friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldVersion, user.FieldWorth:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_best_friend
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				u.Version = int(value.Int64)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldWorth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field worth", values[i])
			} else if value.Valid {
				u.Worth = uint(value.Int64)
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = value.Bool
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_best_friend", value)
			} else if value.Valid {
				u.user_best_friend = new(int)
				*u.user_best_friend = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryCards queries the "cards" edge of the User entity.
func (u *User) QueryCards() *CardQuery {
	return NewUserClient(u.config).QueryCards(u)
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *PetQuery {
	return NewUserClient(u.config).QueryPets(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return NewUserClient(u.config).QueryFriends(u)
}

// QueryBestFriend queries the "best_friend" edge of the User entity.
func (u *User) QueryBestFriend() *UserQuery {
	return NewUserClient(u.config).QueryBestFriend(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("worth=")
	builder.WriteString(fmt.Sprintf("%v", u.Worth))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
