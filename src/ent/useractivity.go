// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
)

// UserActivity is the model entity for the UserActivity schema.
type UserActivity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon *string `json:"icon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserActivityQuery when eager-loading is set.
	Edges           UserActivityEdges `json:"edges"`
	user_activities *int
}

// UserActivityEdges holds the relations/edges for other nodes in the graph.
type UserActivityEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserActivityEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserActivity) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case useractivity.FieldID:
			values[i] = new(sql.NullInt64)
		case useractivity.FieldName, useractivity.FieldURL, useractivity.FieldIcon:
			values[i] = new(sql.NullString)
		case useractivity.FieldCreateTime, useractivity.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case useractivity.ForeignKeys[0]: // user_activities
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserActivity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserActivity fields.
func (ua *UserActivity) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case useractivity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ua.ID = int(value.Int64)
		case useractivity.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ua.CreateTime = value.Time
			}
		case useractivity.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ua.UpdateTime = value.Time
			}
		case useractivity.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ua.Name = value.String
			}
		case useractivity.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				ua.URL = new(string)
				*ua.URL = value.String
			}
		case useractivity.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				ua.Icon = new(string)
				*ua.Icon = value.String
			}
		case useractivity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_activities", value)
			} else if value.Valid {
				ua.user_activities = new(int)
				*ua.user_activities = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserActivity entity.
func (ua *UserActivity) QueryUser() *UserQuery {
	return (&UserActivityClient{config: ua.config}).QueryUser(ua)
}

// Update returns a builder for updating this UserActivity.
// Note that you need to call UserActivity.Unwrap() before calling this method if this UserActivity
// was returned from a transaction, and the transaction was committed or rolled back.
func (ua *UserActivity) Update() *UserActivityUpdateOne {
	return (&UserActivityClient{config: ua.config}).UpdateOne(ua)
}

// Unwrap unwraps the UserActivity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ua *UserActivity) Unwrap() *UserActivity {
	_tx, ok := ua.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserActivity is not a transactional entity")
	}
	ua.config.driver = _tx.drv
	return ua
}

// String implements the fmt.Stringer.
func (ua *UserActivity) String() string {
	var builder strings.Builder
	builder.WriteString("UserActivity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ua.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ua.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ua.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ua.Name)
	builder.WriteString(", ")
	if v := ua.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ua.Icon; v != nil {
		builder.WriteString("icon=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// UserActivities is a parsable slice of UserActivity.
type UserActivities []*UserActivity

func (ua UserActivities) config(cfg config) {
	for _i := range ua {
		ua[_i].config = cfg
	}
}
