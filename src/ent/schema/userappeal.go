package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	userAppealsRef = "appeals"
)

// UserAppeal holds the schema definition for the UserAppeal entity.
type UserAppeal struct {
	ent.Schema
}

// Fields of the UserAppeal.
func (UserAppeal) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty(),
	}
}

func (UserAppeal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the UserAppeal.
func (UserAppeal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdgeName, User.Type).Ref(userAppealsRef).Unique().Required(),
	}
}
