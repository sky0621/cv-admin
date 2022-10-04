// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
)

// UserCareer is the model entity for the UserCareer schema.
type UserCareer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserCareerQuery when eager-loading is set.
	Edges           UserCareerEdges `json:"edges"`
	career_group_id *int
}

// UserCareerEdges holds the relations/edges for other nodes in the graph.
type UserCareerEdges struct {
	// CareerGroup holds the value of the careerGroup edge.
	CareerGroup *UserCareerGroup `json:"careerGroup,omitempty"`
	// CareerDescriptions holds the value of the careerDescriptions edge.
	CareerDescriptions []*UserCareerDescription `json:"careerDescriptions,omitempty"`
	// CareerTasks holds the value of the careerTasks edge.
	CareerTasks []*CareerTask `json:"careerTasks,omitempty"`
	// CareerSkillGroups holds the value of the careerSkillGroups edge.
	CareerSkillGroups []*CareerSkillGroup `json:"careerSkillGroups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CareerGroupOrErr returns the CareerGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserCareerEdges) CareerGroupOrErr() (*UserCareerGroup, error) {
	if e.loadedTypes[0] {
		if e.CareerGroup == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: usercareergroup.Label}
		}
		return e.CareerGroup, nil
	}
	return nil, &NotLoadedError{edge: "careerGroup"}
}

// CareerDescriptionsOrErr returns the CareerDescriptions value or an error if the edge
// was not loaded in eager-loading.
func (e UserCareerEdges) CareerDescriptionsOrErr() ([]*UserCareerDescription, error) {
	if e.loadedTypes[1] {
		return e.CareerDescriptions, nil
	}
	return nil, &NotLoadedError{edge: "careerDescriptions"}
}

// CareerTasksOrErr returns the CareerTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserCareerEdges) CareerTasksOrErr() ([]*CareerTask, error) {
	if e.loadedTypes[2] {
		return e.CareerTasks, nil
	}
	return nil, &NotLoadedError{edge: "careerTasks"}
}

// CareerSkillGroupsOrErr returns the CareerSkillGroups value or an error if the edge
// was not loaded in eager-loading.
func (e UserCareerEdges) CareerSkillGroupsOrErr() ([]*CareerSkillGroup, error) {
	if e.loadedTypes[3] {
		return e.CareerSkillGroups, nil
	}
	return nil, &NotLoadedError{edge: "careerSkillGroups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCareer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercareer.FieldID:
			values[i] = new(sql.NullInt64)
		case usercareer.FieldName, usercareer.FieldFrom, usercareer.FieldTo:
			values[i] = new(sql.NullString)
		case usercareer.FieldCreateTime, usercareer.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case usercareer.ForeignKeys[0]: // career_group_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserCareer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCareer fields.
func (uc *UserCareer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercareer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int(value.Int64)
		case usercareer.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uc.CreateTime = value.Time
			}
		case usercareer.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uc.UpdateTime = value.Time
			}
		case usercareer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				uc.Name = value.String
			}
		case usercareer.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				uc.From = value.String
			}
		case usercareer.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				uc.To = value.String
			}
		case usercareer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field career_group_id", value)
			} else if value.Valid {
				uc.career_group_id = new(int)
				*uc.career_group_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCareerGroup queries the "careerGroup" edge of the UserCareer entity.
func (uc *UserCareer) QueryCareerGroup() *UserCareerGroupQuery {
	return (&UserCareerClient{config: uc.config}).QueryCareerGroup(uc)
}

// QueryCareerDescriptions queries the "careerDescriptions" edge of the UserCareer entity.
func (uc *UserCareer) QueryCareerDescriptions() *UserCareerDescriptionQuery {
	return (&UserCareerClient{config: uc.config}).QueryCareerDescriptions(uc)
}

// QueryCareerTasks queries the "careerTasks" edge of the UserCareer entity.
func (uc *UserCareer) QueryCareerTasks() *CareerTaskQuery {
	return (&UserCareerClient{config: uc.config}).QueryCareerTasks(uc)
}

// QueryCareerSkillGroups queries the "careerSkillGroups" edge of the UserCareer entity.
func (uc *UserCareer) QueryCareerSkillGroups() *CareerSkillGroupQuery {
	return (&UserCareerClient{config: uc.config}).QueryCareerSkillGroups(uc)
}

// Update returns a builder for updating this UserCareer.
// Note that you need to call UserCareer.Unwrap() before calling this method if this UserCareer
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserCareer) Update() *UserCareerUpdateOne {
	return (&UserCareerClient{config: uc.config}).UpdateOne(uc)
}

// Unwrap unwraps the UserCareer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserCareer) Unwrap() *UserCareer {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserCareer is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserCareer) String() string {
	var builder strings.Builder
	builder.WriteString("UserCareer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(uc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(uc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(uc.Name)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(uc.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(uc.To)
	builder.WriteByte(')')
	return builder.String()
}

// UserCareers is a parsable slice of UserCareer.
type UserCareers []*UserCareer

func (uc UserCareers) config(cfg config) {
	for _i := range uc {
		uc[_i].config = cfg
	}
}
