// Code generated by ent, DO NOT EDIT.

package skilltag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldName, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldOrder, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldContainsFold(FieldName, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.SkillTag {
	return predicate.SkillTag(sql.FieldLTE(FieldOrder, v))
}

// HasSkills applies the HasEdge predicate on the "skills" edge.
func HasSkills() predicate.SkillTag {
	return predicate.SkillTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SkillsTable, SkillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillsWith applies the HasEdge predicate on the "skills" edge with a given conditions (other predicates).
func HasSkillsWith(preds ...predicate.Skill) predicate.SkillTag {
	return predicate.SkillTag(func(s *sql.Selector) {
		step := newSkillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SkillTag) predicate.SkillTag {
	return predicate.SkillTag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SkillTag) predicate.SkillTag {
	return predicate.SkillTag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SkillTag) predicate.SkillTag {
	return predicate.SkillTag(sql.NotPredicates(p))
}
