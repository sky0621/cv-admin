package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CareerTask holds the schema definition for the CareerTask entity.
type CareerTask struct {
	ent.Schema
}

// Fields of the CareerTask.
func (CareerTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Validate(maxRuneCount(100)),
	}
}

func (CareerTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the CareerTask.
func (CareerTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("career", UserCareer.Type).Ref("careertasks").Unique().Required(),
		edge.To("careertaskdescriptions", CareerTaskDescription.Type).StorageKey(edge.Column("careertask_id")),
	}
}
