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
		field.String("from").NotEmpty().Validate(fixedRuneCount(6)),
		field.String("to").NotEmpty().Validate(fixedRuneCount(6)),
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
		edge.To("careerdescriptions", UserCareerDescription.Type).StorageKey(edge.Column("career_id")),
		edge.To("careertasks", CareerTask.Type).StorageKey(edge.Column("career_id")),
	}
}
