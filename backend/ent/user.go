// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/F12aPPy/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// USEREMAIL holds the value of the "USER_EMAIL" field.
	USEREMAIL string `json:"USER_EMAIL,omitempty"`
	// USERNAME holds the value of the "USER_NAME" field.
	USERNAME string `json:"USER_NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// CARE holds the value of the CARE edge.
	CARE []*Antenatal
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CAREOrErr returns the CARE value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CAREOrErr() ([]*Antenatal, error) {
	if e.loadedTypes[0] {
		return e.CARE, nil
	}
	return nil, &NotLoadedError{edge: "CARE"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // USER_EMAIL
		&sql.NullString{}, // USER_NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USER_EMAIL", values[0])
	} else if value.Valid {
		u.USEREMAIL = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USER_NAME", values[1])
	} else if value.Valid {
		u.USERNAME = value.String
	}
	return nil
}

// QueryCARE queries the CARE edge of the User.
func (u *User) QueryCARE() *AntenatalQuery {
	return (&UserClient{config: u.config}).QueryCARE(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", USER_EMAIL=")
	builder.WriteString(u.USEREMAIL)
	builder.WriteString(", USER_NAME=")
	builder.WriteString(u.USERNAME)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
