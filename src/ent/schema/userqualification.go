package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const userQualificationsRef = "qualifications"

// UserQualification holds the schema definition for the UserQualification entity.
type UserQualification struct {
	ent.Schema
}

// Fields of the UserQualification.
func (UserQualification) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(80)),
		field.String("organization").Validate(maxRuneCount(80)).Optional().Nillable(),
		field.String("url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
		field.String("got_date").Validate(maxRuneCount(10)).Optional().Nillable(), // FIXME: custom validation
		field.String("memo").Validate(maxRuneCount(400)).Optional().Nillable(),
	}
}

func (UserQualification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserQualification.
func (UserQualification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userQualificationsRef).Unique().Required(),
	}
}
