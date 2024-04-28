// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
)

// UserCareerDescription is the model entity for the UserCareerDescription schema.
type UserCareerDescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserCareerDescriptionQuery when eager-loading is set.
	Edges        UserCareerDescriptionEdges `json:"edges"`
	career_id    *int
	selectValues sql.SelectValues
}

// UserCareerDescriptionEdges holds the relations/edges for other nodes in the graph.
type UserCareerDescriptionEdges struct {
	// Career holds the value of the career edge.
	Career *UserCareer `json:"career,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CareerOrErr returns the Career value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserCareerDescriptionEdges) CareerOrErr() (*UserCareer, error) {
	if e.Career != nil {
		return e.Career, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: usercareer.Label}
	}
	return nil, &NotLoadedError{edge: "career"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCareerDescription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercareerdescription.FieldID:
			values[i] = new(sql.NullInt64)
		case usercareerdescription.FieldDescription:
			values[i] = new(sql.NullString)
		case usercareerdescription.ForeignKeys[0]: // career_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCareerDescription fields.
func (ucd *UserCareerDescription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercareerdescription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ucd.ID = int(value.Int64)
		case usercareerdescription.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ucd.Description = value.String
			}
		case usercareerdescription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field career_id", value)
			} else if value.Valid {
				ucd.career_id = new(int)
				*ucd.career_id = int(value.Int64)
			}
		default:
			ucd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserCareerDescription.
// This includes values selected through modifiers, order, etc.
func (ucd *UserCareerDescription) Value(name string) (ent.Value, error) {
	return ucd.selectValues.Get(name)
}

// QueryCareer queries the "career" edge of the UserCareerDescription entity.
func (ucd *UserCareerDescription) QueryCareer() *UserCareerQuery {
	return NewUserCareerDescriptionClient(ucd.config).QueryCareer(ucd)
}

// Update returns a builder for updating this UserCareerDescription.
// Note that you need to call UserCareerDescription.Unwrap() before calling this method if this UserCareerDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (ucd *UserCareerDescription) Update() *UserCareerDescriptionUpdateOne {
	return NewUserCareerDescriptionClient(ucd.config).UpdateOne(ucd)
}

// Unwrap unwraps the UserCareerDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ucd *UserCareerDescription) Unwrap() *UserCareerDescription {
	_tx, ok := ucd.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserCareerDescription is not a transactional entity")
	}
	ucd.config.driver = _tx.drv
	return ucd
}

// String implements the fmt.Stringer.
func (ucd *UserCareerDescription) String() string {
	var builder strings.Builder
	builder.WriteString("UserCareerDescription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ucd.ID))
	builder.WriteString("description=")
	builder.WriteString(ucd.Description)
	builder.WriteByte(')')
	return builder.String()
}

// UserCareerDescriptions is a parsable slice of UserCareerDescription.
type UserCareerDescriptions []*UserCareerDescription
