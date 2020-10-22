// Code generated by entc, DO NOT EDIT.

package room

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/moominzie/user-record/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Rname applies equality check predicate on the "rname" field. It's identical to RnameEQ.
func Rname(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRname), v))
	})
}

// RnameEQ applies the EQ predicate on the "rname" field.
func RnameEQ(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRname), v))
	})
}

// RnameNEQ applies the NEQ predicate on the "rname" field.
func RnameNEQ(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRname), v))
	})
}

// RnameIn applies the In predicate on the "rname" field.
func RnameIn(vs ...string) predicate.Room {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Room(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRname), v...))
	})
}

// RnameNotIn applies the NotIn predicate on the "rname" field.
func RnameNotIn(vs ...string) predicate.Room {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Room(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRname), v...))
	})
}

// RnameGT applies the GT predicate on the "rname" field.
func RnameGT(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRname), v))
	})
}

// RnameGTE applies the GTE predicate on the "rname" field.
func RnameGTE(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRname), v))
	})
}

// RnameLT applies the LT predicate on the "rname" field.
func RnameLT(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRname), v))
	})
}

// RnameLTE applies the LTE predicate on the "rname" field.
func RnameLTE(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRname), v))
	})
}

// RnameContains applies the Contains predicate on the "rname" field.
func RnameContains(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRname), v))
	})
}

// RnameHasPrefix applies the HasPrefix predicate on the "rname" field.
func RnameHasPrefix(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRname), v))
	})
}

// RnameHasSuffix applies the HasSuffix predicate on the "rname" field.
func RnameHasSuffix(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRname), v))
	})
}

// RnameEqualFold applies the EqualFold predicate on the "rname" field.
func RnameEqualFold(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRname), v))
	})
}

// RnameContainsFold applies the ContainsFold predicate on the "rname" field.
func RnameContainsFold(v string) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRname), v))
	})
}

// HasBuilding applies the HasEdge predicate on the "building" edge.
func HasBuilding() predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildingTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BuildingTable, BuildingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildingWith applies the HasEdge predicate on the "building" edge with a given conditions (other predicates).
func HasBuildingWith(preds ...predicate.Building) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BuildingTable, BuildingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserInformations applies the HasEdge predicate on the "user_informations" edge.
func HasUserInformations() predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInformationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserInformationsTable, UserInformationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserInformationsWith applies the HasEdge predicate on the "user_informations" edge with a given conditions (other predicates).
func HasUserInformationsWith(preds ...predicate.User) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInformationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserInformationsTable, UserInformationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
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
func Not(p predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		p(s.Not())
	})
}
