// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/skilltag"
)

// Skill is the model entity for the Skill schema.
type Skill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkillQuery when eager-loading is set.
	Edges        SkillEdges `json:"edges"`
	tag_id       *int
	selectValues sql.SelectValues
}

// SkillEdges holds the relations/edges for other nodes in the graph.
type SkillEdges struct {
	// SkillTag holds the value of the skillTag edge.
	SkillTag *SkillTag `json:"skillTag,omitempty"`
	// CareerSkills holds the value of the careerSkills edge.
	CareerSkills []*CareerSkill `json:"careerSkills,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SkillTagOrErr returns the SkillTag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkillEdges) SkillTagOrErr() (*SkillTag, error) {
	if e.loadedTypes[0] {
		if e.SkillTag == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: skilltag.Label}
		}
		return e.SkillTag, nil
	}
	return nil, &NotLoadedError{edge: "skillTag"}
}

// CareerSkillsOrErr returns the CareerSkills value or an error if the edge
// was not loaded in eager-loading.
func (e SkillEdges) CareerSkillsOrErr() ([]*CareerSkill, error) {
	if e.loadedTypes[1] {
		return e.CareerSkills, nil
	}
	return nil, &NotLoadedError{edge: "careerSkills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Skill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case skill.FieldID:
			values[i] = new(sql.NullInt64)
		case skill.FieldName, skill.FieldCode, skill.FieldURL:
			values[i] = new(sql.NullString)
		case skill.FieldCreateTime, skill.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case skill.ForeignKeys[0]: // tag_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Skill fields.
func (s *Skill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case skill.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case skill.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case skill.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case skill.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case skill.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = new(string)
				*s.URL = value.String
			}
		case skill.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tag_id", value)
			} else if value.Valid {
				s.tag_id = new(int)
				*s.tag_id = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Skill.
// This includes values selected through modifiers, order, etc.
func (s *Skill) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QuerySkillTag queries the "skillTag" edge of the Skill entity.
func (s *Skill) QuerySkillTag() *SkillTagQuery {
	return NewSkillClient(s.config).QuerySkillTag(s)
}

// QueryCareerSkills queries the "careerSkills" edge of the Skill entity.
func (s *Skill) QueryCareerSkills() *CareerSkillQuery {
	return NewSkillClient(s.config).QueryCareerSkills(s)
}

// Update returns a builder for updating this Skill.
// Note that you need to call Skill.Unwrap() before calling this method if this Skill
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Skill) Update() *SkillUpdateOne {
	return NewSkillClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Skill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Skill) Unwrap() *Skill {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Skill is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Skill) String() string {
	var builder strings.Builder
	builder.WriteString("Skill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	if v := s.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Skills is a parsable slice of Skill.
type Skills []*Skill
