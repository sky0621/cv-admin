// Code generated by ent, DO NOT EDIT.

package usernoteitem

const (
	// Label holds the string label denoting the usernoteitem type in the database.
	Label = "user_note_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the usernoteitem in the database.
	Table = "user_note_items"
)

// Columns holds all SQL columns for usernoteitem fields.
var Columns = []string{
	FieldID,
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
