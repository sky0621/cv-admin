// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usernote"
)

// UserNote is the model entity for the UserNote schema.
type UserNote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo *string `json:"memo,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserNoteQuery when eager-loading is set.
	Edges        UserNoteEdges `json:"edges"`
	user_id      *int
	selectValues sql.SelectValues
}

// UserNoteEdges holds the relations/edges for other nodes in the graph.
type UserNoteEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// NoteItems holds the value of the noteItems edge.
	NoteItems []*UserNoteItem `json:"noteItems,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserNoteEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// NoteItemsOrErr returns the NoteItems value or an error if the edge
// was not loaded in eager-loading.
func (e UserNoteEdges) NoteItemsOrErr() ([]*UserNoteItem, error) {
	if e.loadedTypes[1] {
		return e.NoteItems, nil
	}
	return nil, &NotLoadedError{edge: "noteItems"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserNote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usernote.FieldID:
			values[i] = new(sql.NullInt64)
		case usernote.FieldLabel, usernote.FieldMemo:
			values[i] = new(sql.NullString)
		case usernote.FieldCreateTime, usernote.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case usernote.ForeignKeys[0]: // user_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserNote fields.
func (un *UserNote) assignValues(columns []string, values []any) error {
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
		case usernote.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				un.CreateTime = value.Time
			}
		case usernote.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				un.UpdateTime = value.Time
			}
		case usernote.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				un.Label = value.String
			}
		case usernote.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				un.Memo = new(string)
				*un.Memo = value.String
			}
		case usernote.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_id", value)
			} else if value.Valid {
				un.user_id = new(int)
				*un.user_id = int(value.Int64)
			}
		default:
			un.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserNote.
// This includes values selected through modifiers, order, etc.
func (un *UserNote) Value(name string) (ent.Value, error) {
	return un.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserNote entity.
func (un *UserNote) QueryUser() *UserQuery {
	return NewUserNoteClient(un.config).QueryUser(un)
}

// QueryNoteItems queries the "noteItems" edge of the UserNote entity.
func (un *UserNote) QueryNoteItems() *UserNoteItemQuery {
	return NewUserNoteClient(un.config).QueryNoteItems(un)
}

// Update returns a builder for updating this UserNote.
// Note that you need to call UserNote.Unwrap() before calling this method if this UserNote
// was returned from a transaction, and the transaction was committed or rolled back.
func (un *UserNote) Update() *UserNoteUpdateOne {
	return NewUserNoteClient(un.config).UpdateOne(un)
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
	builder.WriteString(fmt.Sprintf("id=%v, ", un.ID))
	builder.WriteString("create_time=")
	builder.WriteString(un.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(un.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(un.Label)
	builder.WriteString(", ")
	if v := un.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// UserNotes is a parsable slice of UserNote.
type UserNotes []*UserNote
