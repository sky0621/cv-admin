// Code generated by ent, DO NOT EDIT.

package careertaskdescription

const (
	// Label holds the string label denoting the careertaskdescription type in the database.
	Label = "career_task_description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCareertask holds the string denoting the careertask edge name in mutations.
	EdgeCareertask = "careertask"
	// Table holds the table name of the careertaskdescription in the database.
	Table = "career_task_descriptions"
	// CareertaskTable is the table that holds the careertask relation/edge.
	CareertaskTable = "career_task_descriptions"
	// CareertaskInverseTable is the table name for the CareerTask entity.
	// It exists in this package in order to avoid circular dependency with the "careertask" package.
	CareertaskInverseTable = "career_tasks"
	// CareertaskColumn is the table column denoting the careertask relation/edge.
	CareertaskColumn = "careertask_id"
)

// Columns holds all SQL columns for careertaskdescription fields.
var Columns = []string{
	FieldID,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "career_task_descriptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"careertask_id",
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
