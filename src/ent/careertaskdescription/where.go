// Code generated by ent, DO NOT EDIT.

package careertaskdescription

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldEQ(FieldDescription, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(sql.FieldContainsFold(FieldDescription, v))
}

// HasCareerTask applies the HasEdge predicate on the "careerTask" edge.
func HasCareerTask() predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CareerTaskTable, CareerTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCareerTaskWith applies the HasEdge predicate on the "careerTask" edge with a given conditions (other predicates).
func HasCareerTaskWith(preds ...predicate.CareerTask) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(func(s *sql.Selector) {
		step := newCareerTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CareerTaskDescription) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CareerTaskDescription) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CareerTaskDescription) predicate.CareerTaskDescription {
	return predicate.CareerTaskDescription(func(s *sql.Selector) {
		p(s.Not())
	})
}
