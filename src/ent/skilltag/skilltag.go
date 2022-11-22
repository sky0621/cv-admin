// Code generated by ent, DO NOT EDIT.

package skilltag

const (
	// Label holds the string label denoting the skilltag type in the database.
	Label = "skill_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// Table holds the table name of the skilltag in the database.
	Table = "skill_tags"
)

// Columns holds all SQL columns for skilltag fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldKey,
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
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
)