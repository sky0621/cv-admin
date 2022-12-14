package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const careerTaskDescriptionsRef = "careerTaskDescriptions"

// CareerTaskDescription holds the schema definition for the CareerTaskDescription entity.
type CareerTaskDescription struct {
	ent.Schema
}

// Fields of the CareerTaskDescription.
func (CareerTaskDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty().Validate(maxRuneCount(400)),
	}
}

// Edges of the CareerTaskDescription.
func (CareerTaskDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(careerTaskEdgeName, CareerTask.Type).Ref(careerTaskDescriptionsRef).Unique().Required(),
	}
}
