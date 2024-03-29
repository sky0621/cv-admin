// Code generated by ent, DO NOT EDIT.

package usercareerdescription

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usercareerdescription type in the database.
	Label = "user_career_description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCareer holds the string denoting the career edge name in mutations.
	EdgeCareer = "career"
	// Table holds the table name of the usercareerdescription in the database.
	Table = "user_career_descriptions"
	// CareerTable is the table that holds the career relation/edge.
	CareerTable = "user_career_descriptions"
	// CareerInverseTable is the table name for the UserCareer entity.
	// It exists in this package in order to avoid circular dependency with the "usercareer" package.
	CareerInverseTable = "user_careers"
	// CareerColumn is the table column denoting the career relation/edge.
	CareerColumn = "career_id"
)

// Columns holds all SQL columns for usercareerdescription fields.
var Columns = []string{
	FieldID,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_career_descriptions"
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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

// OrderOption defines the ordering options for the UserCareerDescription queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCareerField orders the results by career field.
func ByCareerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCareerStep(), sql.OrderByField(field, opts...))
	}
}
func newCareerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CareerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CareerTable, CareerColumn),
	)
}
