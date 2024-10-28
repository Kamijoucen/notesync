// Code generated by ent, DO NOT EDIT.

package fileitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kamijoucen/notesync/apps/server/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FileItem {
	return predicate.FileItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FileItem {
	return predicate.FileItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FileItem {
	return predicate.FileItem(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldName, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldHash, v))
}

// IsDir applies equality check predicate on the "is_dir" field. It's identical to IsDirEQ.
func IsDir(v bool) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldIsDir, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.FileItem {
	return predicate.FileItem(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.FileItem {
	return predicate.FileItem(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.FileItem {
	return predicate.FileItem(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldContainsFold(FieldName, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.FileItem {
	return predicate.FileItem(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.FileItem {
	return predicate.FileItem(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.FileItem {
	return predicate.FileItem(sql.FieldContainsFold(FieldHash, v))
}

// IsDirEQ applies the EQ predicate on the "is_dir" field.
func IsDirEQ(v bool) predicate.FileItem {
	return predicate.FileItem(sql.FieldEQ(FieldIsDir, v))
}

// IsDirNEQ applies the NEQ predicate on the "is_dir" field.
func IsDirNEQ(v bool) predicate.FileItem {
	return predicate.FileItem(sql.FieldNEQ(FieldIsDir, v))
}

// HasRepository applies the HasEdge predicate on the "repository" edge.
func HasRepository() predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepositoryWith applies the HasEdge predicate on the "repository" edge with a given conditions (other predicates).
func HasRepositoryWith(preds ...predicate.Repository) predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := newRepositoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.FileItem) predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.FileItem) predicate.FileItem {
	return predicate.FileItem(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FileItem) predicate.FileItem {
	return predicate.FileItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FileItem) predicate.FileItem {
	return predicate.FileItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FileItem) predicate.FileItem {
	return predicate.FileItem(sql.NotPredicates(p))
}