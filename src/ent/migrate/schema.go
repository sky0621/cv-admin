// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CareerSkillsColumns holds the columns for the "career_skills" table.
	CareerSkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CareerSkillsTable holds the schema information for the "career_skills" table.
	CareerSkillsTable = &schema.Table{
		Name:       "career_skills",
		Columns:    CareerSkillsColumns,
		PrimaryKey: []*schema.Column{CareerSkillsColumns[0]},
	}
	// CareerSkillGroupsColumns holds the columns for the "career_skill_groups" table.
	CareerSkillGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CareerSkillGroupsTable holds the schema information for the "career_skill_groups" table.
	CareerSkillGroupsTable = &schema.Table{
		Name:       "career_skill_groups",
		Columns:    CareerSkillGroupsColumns,
		PrimaryKey: []*schema.Column{CareerSkillGroupsColumns[0]},
	}
	// CareerTasksColumns holds the columns for the "career_tasks" table.
	CareerTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CareerTasksTable holds the schema information for the "career_tasks" table.
	CareerTasksTable = &schema.Table{
		Name:       "career_tasks",
		Columns:    CareerTasksColumns,
		PrimaryKey: []*schema.Column{CareerTasksColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "birthday_year", Type: field.TypeInt},
		{Name: "birthday_month", Type: field.TypeInt},
		{Name: "birthday_day", Type: field.TypeInt},
		{Name: "job", Type: field.TypeString, Nullable: true},
		{Name: "belong_to", Type: field.TypeString, Nullable: true},
		{Name: "pr", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserActivitiesColumns holds the columns for the "user_activities" table.
	UserActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserActivitiesTable holds the schema information for the "user_activities" table.
	UserActivitiesTable = &schema.Table{
		Name:       "user_activities",
		Columns:    UserActivitiesColumns,
		PrimaryKey: []*schema.Column{UserActivitiesColumns[0]},
	}
	// UserCareersColumns holds the columns for the "user_careers" table.
	UserCareersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserCareersTable holds the schema information for the "user_careers" table.
	UserCareersTable = &schema.Table{
		Name:       "user_careers",
		Columns:    UserCareersColumns,
		PrimaryKey: []*schema.Column{UserCareersColumns[0]},
	}
	// UserCareerGroupsColumns holds the columns for the "user_career_groups" table.
	UserCareerGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserCareerGroupsTable holds the schema information for the "user_career_groups" table.
	UserCareerGroupsTable = &schema.Table{
		Name:       "user_career_groups",
		Columns:    UserCareerGroupsColumns,
		PrimaryKey: []*schema.Column{UserCareerGroupsColumns[0]},
	}
	// UserNotesColumns holds the columns for the "user_notes" table.
	UserNotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserNotesTable holds the schema information for the "user_notes" table.
	UserNotesTable = &schema.Table{
		Name:       "user_notes",
		Columns:    UserNotesColumns,
		PrimaryKey: []*schema.Column{UserNotesColumns[0]},
	}
	// UserNoteItemsColumns holds the columns for the "user_note_items" table.
	UserNoteItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserNoteItemsTable holds the schema information for the "user_note_items" table.
	UserNoteItemsTable = &schema.Table{
		Name:       "user_note_items",
		Columns:    UserNoteItemsColumns,
		PrimaryKey: []*schema.Column{UserNoteItemsColumns[0]},
	}
	// UserQualificationsColumns holds the columns for the "user_qualifications" table.
	UserQualificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UserQualificationsTable holds the schema information for the "user_qualifications" table.
	UserQualificationsTable = &schema.Table{
		Name:       "user_qualifications",
		Columns:    UserQualificationsColumns,
		PrimaryKey: []*schema.Column{UserQualificationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CareerSkillsTable,
		CareerSkillGroupsTable,
		CareerTasksTable,
		UsersTable,
		UserActivitiesTable,
		UserCareersTable,
		UserCareerGroupsTable,
		UserNotesTable,
		UserNoteItemsTable,
		UserQualificationsTable,
	}
)

func init() {
}