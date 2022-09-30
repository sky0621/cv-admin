// Code generated by ent, DO NOT EDIT.

package userqualification

import (
	"time"
)

const (
	// Label holds the string label denoting the userqualification type in the database.
	Label = "user_qualification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldGotDate holds the string denoting the got_date field in the database.
	FieldGotDate = "got_date"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userqualification in the database.
	Table = "user_qualifications"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_qualifications"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userqualification fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldOrganization,
	FieldURL,
	FieldGotDate,
	FieldMemo,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_qualifications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
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
	// OrganizationValidator is a validator for the "organization" field. It is called by the builders before save.
	OrganizationValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// GotDateValidator is a validator for the "got_date" field. It is called by the builders before save.
	GotDateValidator func(string) error
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
)
