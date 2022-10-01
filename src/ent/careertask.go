// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
)

// CareerTask is the model entity for the CareerTask schema.
type CareerTask struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CareerTaskQuery when eager-loading is set.
	Edges     CareerTaskEdges `json:"edges"`
	career_id *int
}

// CareerTaskEdges holds the relations/edges for other nodes in the graph.
type CareerTaskEdges struct {
	// Career holds the value of the career edge.
	Career *UserCareer `json:"career,omitempty"`
	// Careertaskdescriptions holds the value of the careertaskdescriptions edge.
	Careertaskdescriptions []*CareerTaskDescription `json:"careertaskdescriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CareerOrErr returns the Career value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CareerTaskEdges) CareerOrErr() (*UserCareer, error) {
	if e.loadedTypes[0] {
		if e.Career == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: usercareer.Label}
		}
		return e.Career, nil
	}
	return nil, &NotLoadedError{edge: "career"}
}

// CareertaskdescriptionsOrErr returns the Careertaskdescriptions value or an error if the edge
// was not loaded in eager-loading.
func (e CareerTaskEdges) CareertaskdescriptionsOrErr() ([]*CareerTaskDescription, error) {
	if e.loadedTypes[1] {
		return e.Careertaskdescriptions, nil
	}
	return nil, &NotLoadedError{edge: "careertaskdescriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CareerTask) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case careertask.FieldID:
			values[i] = new(sql.NullInt64)
		case careertask.FieldName:
			values[i] = new(sql.NullString)
		case careertask.FieldCreateTime, careertask.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case careertask.ForeignKeys[0]: // career_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CareerTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CareerTask fields.
func (ct *CareerTask) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case careertask.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case careertask.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ct.CreateTime = value.Time
			}
		case careertask.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ct.UpdateTime = value.Time
			}
		case careertask.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ct.Name = value.String
			}
		case careertask.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field career_id", value)
			} else if value.Valid {
				ct.career_id = new(int)
				*ct.career_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCareer queries the "career" edge of the CareerTask entity.
func (ct *CareerTask) QueryCareer() *UserCareerQuery {
	return (&CareerTaskClient{config: ct.config}).QueryCareer(ct)
}

// QueryCareertaskdescriptions queries the "careertaskdescriptions" edge of the CareerTask entity.
func (ct *CareerTask) QueryCareertaskdescriptions() *CareerTaskDescriptionQuery {
	return (&CareerTaskClient{config: ct.config}).QueryCareertaskdescriptions(ct)
}

// Update returns a builder for updating this CareerTask.
// Note that you need to call CareerTask.Unwrap() before calling this method if this CareerTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CareerTask) Update() *CareerTaskUpdateOne {
	return (&CareerTaskClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the CareerTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CareerTask) Unwrap() *CareerTask {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: CareerTask is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CareerTask) String() string {
	var builder strings.Builder
	builder.WriteString("CareerTask(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ct.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ct.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ct.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CareerTasks is a parsable slice of CareerTask.
type CareerTasks []*CareerTask

func (ct CareerTasks) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
