// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldBirthdayYear holds the string denoting the birthday_year field in the database.
	FieldBirthdayYear = "birthday_year"
	// FieldBirthdayMonth holds the string denoting the birthday_month field in the database.
	FieldBirthdayMonth = "birthday_month"
	// FieldBirthdayDay holds the string denoting the birthday_day field in the database.
	FieldBirthdayDay = "birthday_day"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// FieldBelongTo holds the string denoting the belong_to field in the database.
	FieldBelongTo = "belong_to"
	// FieldPr holds the string denoting the pr field in the database.
	FieldPr = "pr"
	// EdgeActivities holds the string denoting the activities edge name in mutations.
	EdgeActivities = "activities"
	// EdgeQualifications holds the string denoting the qualifications edge name in mutations.
	EdgeQualifications = "qualifications"
	// EdgeCareerGroups holds the string denoting the careergroups edge name in mutations.
	EdgeCareerGroups = "careerGroups"
	// EdgeNotes holds the string denoting the notes edge name in mutations.
	EdgeNotes = "notes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ActivitiesTable is the table that holds the activities relation/edge.
	ActivitiesTable = "user_activities"
	// ActivitiesInverseTable is the table name for the UserActivity entity.
	// It exists in this package in order to avoid circular dependency with the "useractivity" package.
	ActivitiesInverseTable = "user_activities"
	// ActivitiesColumn is the table column denoting the activities relation/edge.
	ActivitiesColumn = "user_id"
	// QualificationsTable is the table that holds the qualifications relation/edge.
	QualificationsTable = "user_qualifications"
	// QualificationsInverseTable is the table name for the UserQualification entity.
	// It exists in this package in order to avoid circular dependency with the "userqualification" package.
	QualificationsInverseTable = "user_qualifications"
	// QualificationsColumn is the table column denoting the qualifications relation/edge.
	QualificationsColumn = "user_id"
	// CareerGroupsTable is the table that holds the careerGroups relation/edge.
	CareerGroupsTable = "user_career_groups"
	// CareerGroupsInverseTable is the table name for the UserCareerGroup entity.
	// It exists in this package in order to avoid circular dependency with the "usercareergroup" package.
	CareerGroupsInverseTable = "user_career_groups"
	// CareerGroupsColumn is the table column denoting the careerGroups relation/edge.
	CareerGroupsColumn = "user_id"
	// NotesTable is the table that holds the notes relation/edge.
	NotesTable = "user_notes"
	// NotesInverseTable is the table name for the UserNote entity.
	// It exists in this package in order to avoid circular dependency with the "usernote" package.
	NotesInverseTable = "user_notes"
	// NotesColumn is the table column denoting the notes relation/edge.
	NotesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldNickname,
	FieldAvatarURL,
	FieldBirthdayYear,
	FieldBirthdayMonth,
	FieldBirthdayDay,
	FieldJob,
	FieldBelongTo,
	FieldPr,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	AvatarURLValidator func(string) error
	// BirthdayYearValidator is a validator for the "birthday_year" field. It is called by the builders before save.
	BirthdayYearValidator func(int) error
	// BirthdayMonthValidator is a validator for the "birthday_month" field. It is called by the builders before save.
	BirthdayMonthValidator func(int) error
	// BirthdayDayValidator is a validator for the "birthday_day" field. It is called by the builders before save.
	BirthdayDayValidator func(int) error
	// JobValidator is a validator for the "job" field. It is called by the builders before save.
	JobValidator func(string) error
	// BelongToValidator is a validator for the "belong_to" field. It is called by the builders before save.
	BelongToValidator func(string) error
	// PrValidator is a validator for the "pr" field. It is called by the builders before save.
	PrValidator func(string) error
)
