// Code generated by ent, DO NOT EDIT.

package skilltag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the skilltag type in the database.
	Label = "skill_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSkills holds the string denoting the skills edge name in mutations.
	EdgeSkills = "skills"
	// Table holds the table name of the skilltag in the database.
	Table = "skill_tags"
	// SkillsTable is the table that holds the skills relation/edge.
	SkillsTable = "skills"
	// SkillsInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillsInverseTable = "skills"
	// SkillsColumn is the table column denoting the skills relation/edge.
	SkillsColumn = "tag_id"
)

// Columns holds all SQL columns for skilltag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the SkillTag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySkillsCount orders the results by skills count.
func BySkillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSkillsStep(), opts...)
	}
}

// BySkills orders the results by skills terms.
func BySkills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSkillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSkillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SkillsTable, SkillsColumn),
	)
}
