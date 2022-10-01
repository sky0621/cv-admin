// Code generated by ent, DO NOT EDIT.

package careerskillgroup

import (
	"time"
)

const (
	// Label holds the string label denoting the careerskillgroup type in the database.
	Label = "career_skill_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldLabel holds the string denoting the label field in the database.
	FieldLabel = "label"
	// EdgeCareer holds the string denoting the career edge name in mutations.
	EdgeCareer = "career"
	// Table holds the table name of the careerskillgroup in the database.
	Table = "career_skill_groups"
	// CareerTable is the table that holds the career relation/edge.
	CareerTable = "career_skill_groups"
	// CareerInverseTable is the table name for the UserCareer entity.
	// It exists in this package in order to avoid circular dependency with the "usercareer" package.
	CareerInverseTable = "user_careers"
	// CareerColumn is the table column denoting the career relation/edge.
	CareerColumn = "career_id"
)

// Columns holds all SQL columns for careerskillgroup fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldLabel,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "career_skill_groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"career_id",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// LabelValidator is a validator for the "label" field. It is called by the builders before save.
	LabelValidator func(string) error
)
