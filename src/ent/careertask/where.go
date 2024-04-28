// Code generated by ent, DO NOT EDIT.

package careertask

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sky0621/cv-admin/src/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldName, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CareerTask {
	return predicate.CareerTask(sql.FieldContainsFold(FieldName, v))
}

// HasCareer applies the HasEdge predicate on the "career" edge.
func HasCareer() predicate.CareerTask {
	return predicate.CareerTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CareerTable, CareerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCareerWith applies the HasEdge predicate on the "career" edge with a given conditions (other predicates).
func HasCareerWith(preds ...predicate.UserCareer) predicate.CareerTask {
	return predicate.CareerTask(func(s *sql.Selector) {
		step := newCareerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCareerTaskDescriptions applies the HasEdge predicate on the "careerTaskDescriptions" edge.
func HasCareerTaskDescriptions() predicate.CareerTask {
	return predicate.CareerTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CareerTaskDescriptionsTable, CareerTaskDescriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCareerTaskDescriptionsWith applies the HasEdge predicate on the "careerTaskDescriptions" edge with a given conditions (other predicates).
func HasCareerTaskDescriptionsWith(preds ...predicate.CareerTaskDescription) predicate.CareerTask {
	return predicate.CareerTask(func(s *sql.Selector) {
		step := newCareerTaskDescriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CareerTask) predicate.CareerTask {
	return predicate.CareerTask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CareerTask) predicate.CareerTask {
	return predicate.CareerTask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CareerTask) predicate.CareerTask {
	return predicate.CareerTask(sql.NotPredicates(p))
}
