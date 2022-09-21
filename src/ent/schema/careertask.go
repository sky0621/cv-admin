package schema

import "entgo.io/ent"

// CareerTask holds the schema definition for the CareerTask entity.
type CareerTask struct {
	ent.Schema
}

// Fields of the CareerTask.
func (CareerTask) Fields() []ent.Field {
	return nil
}

func (CareerTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the CareerTask.
func (CareerTask) Edges() []ent.Edge {
	return nil
}
