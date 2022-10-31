package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const userCareerDescriptionsRef = "careerDescriptions"

// UserCareerDescription holds the schema definition for the UserCareerDescription entity.
type UserCareerDescription struct {
	ent.Schema
}

// Fields of the UserCareerDescription.
func (UserCareerDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty().Validate(maxRuneCount(400)),
	}
}

// Edges of the UserCareerDescription.
func (UserCareerDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userCareerEdgeName, UserCareer.Type).Ref(userCareerDescriptionsRef).Unique().Required(),
	}
}
