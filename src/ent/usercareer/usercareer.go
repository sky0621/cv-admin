// Code generated by ent, DO NOT EDIT.

package usercareer

import (
	"time"
)

const (
	// Label holds the string label denoting the usercareer type in the database.
	Label = "user_career"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// EdgeCareergroup holds the string denoting the careergroup edge name in mutations.
	EdgeCareergroup = "careergroup"
	// EdgeCareerdescriptions holds the string denoting the careerdescriptions edge name in mutations.
	EdgeCareerdescriptions = "careerdescriptions"
	// EdgeCareertasks holds the string denoting the careertasks edge name in mutations.
	EdgeCareertasks = "careertasks"
	// Table holds the table name of the usercareer in the database.
	Table = "user_careers"
	// CareergroupTable is the table that holds the careergroup relation/edge.
	CareergroupTable = "user_careers"
	// CareergroupInverseTable is the table name for the UserCareerGroup entity.
	// It exists in this package in order to avoid circular dependency with the "usercareergroup" package.
	CareergroupInverseTable = "user_career_groups"
	// CareergroupColumn is the table column denoting the careergroup relation/edge.
	CareergroupColumn = "careergroup_id"
	// CareerdescriptionsTable is the table that holds the careerdescriptions relation/edge.
	CareerdescriptionsTable = "user_career_descriptions"
	// CareerdescriptionsInverseTable is the table name for the UserCareerDescription entity.
	// It exists in this package in order to avoid circular dependency with the "usercareerdescription" package.
	CareerdescriptionsInverseTable = "user_career_descriptions"
	// CareerdescriptionsColumn is the table column denoting the careerdescriptions relation/edge.
	CareerdescriptionsColumn = "career_id"
	// CareertasksTable is the table that holds the careertasks relation/edge.
	CareertasksTable = "career_tasks"
	// CareertasksInverseTable is the table name for the CareerTask entity.
	// It exists in this package in order to avoid circular dependency with the "careertask" package.
	CareertasksInverseTable = "career_tasks"
	// CareertasksColumn is the table column denoting the careertasks relation/edge.
	CareertasksColumn = "career_id"
)

// Columns holds all SQL columns for usercareer fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldFrom,
	FieldTo,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_careers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"careergroup_id",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// FromValidator is a validator for the "from" field. It is called by the builders before save.
	FromValidator func(string) error
	// ToValidator is a validator for the "to" field. It is called by the builders before save.
	ToValidator func(string) error
)
