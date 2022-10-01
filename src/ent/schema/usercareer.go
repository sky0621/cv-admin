package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCareer holds the schema definition for the UserCareer entity.
type UserCareer struct {
	ent.Schema
}

// Fields of the UserCareer.
func (UserCareer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)),
		field.String("url").Validate(rangeRuneCount(1, 4096)).Optional().Nillable(),
		field.String("icon").Validate(maxRuneCount(40)).Optional().Nillable(),
	}
}

func (UserCareer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserCareer.
func (UserCareer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("careergroup", UserCareerGroup.Type).Ref("careers").Unique().Required(),
	}
}
