package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	userSolutionsRef = "solutions"
)

// UserSolution holds the schema definition for the UserSolution entity.
type UserSolution struct {
	ent.Schema
}

// Fields of the UserSolution.
func (UserSolution) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty(),
	}
}

func (UserSolution) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserSolution.
func (UserSolution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userSolutionsRef).Unique().Required(),
	}
}
