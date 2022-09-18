// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/usernote"
)

// UserNote is the model entity for the UserNote schema.
type UserNote struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserNote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case usernote.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserNote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserNote fields.
func (un *UserNote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usernote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			un.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this UserNote.
// Note that you need to call UserNote.Unwrap() before calling this method if this UserNote
// was returned from a transaction, and the transaction was committed or rolled back.
func (un *UserNote) Update() *UserNoteUpdateOne {
	return (&UserNoteClient{config: un.config}).UpdateOne(un)
}

// Unwrap unwraps the UserNote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (un *UserNote) Unwrap() *UserNote {
	_tx, ok := un.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserNote is not a transactional entity")
	}
	un.config.driver = _tx.drv
	return un
}

// String implements the fmt.Stringer.
func (un *UserNote) String() string {
	var builder strings.Builder
	builder.WriteString("UserNote(")
	builder.WriteString(fmt.Sprintf("id=%v", un.ID))
	builder.WriteByte(')')
	return builder.String()
}

// UserNotes is a parsable slice of UserNote.
type UserNotes []*UserNote

func (un UserNotes) config(cfg config) {
	for _i := range un {
		un[_i].config = cfg
	}
}