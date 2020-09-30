// Code generated by entc, DO NOT EDIT.

package antenatal

import (
	"time"

	"github.com/F12aPPy/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ADDEDTIME applies equality check predicate on the "ADDED_TIME" field. It's identical to ADDEDTIMEEQ.
func ADDEDTIME(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMEEQ applies the EQ predicate on the "ADDED_TIME" field.
func ADDEDTIMEEQ(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMENEQ applies the NEQ predicate on the "ADDED_TIME" field.
func ADDEDTIMENEQ(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMEIn applies the In predicate on the "ADDED_TIME" field.
func ADDEDTIMEIn(vs ...time.Time) predicate.Antenatal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Antenatal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldADDEDTIME), v...))
	})
}

// ADDEDTIMENotIn applies the NotIn predicate on the "ADDED_TIME" field.
func ADDEDTIMENotIn(vs ...time.Time) predicate.Antenatal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Antenatal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldADDEDTIME), v...))
	})
}

// ADDEDTIMEGT applies the GT predicate on the "ADDED_TIME" field.
func ADDEDTIMEGT(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMEGTE applies the GTE predicate on the "ADDED_TIME" field.
func ADDEDTIMEGTE(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMELT applies the LT predicate on the "ADDED_TIME" field.
func ADDEDTIMELT(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldADDEDTIME), v))
	})
}

// ADDEDTIMELTE applies the LTE predicate on the "ADDED_TIME" field.
func ADDEDTIMELTE(v time.Time) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldADDEDTIME), v))
	})
}

// HasGETMOM applies the HasEdge predicate on the "GETMOM" edge.
func HasGETMOM() predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GETMOMTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GETMOMTable, GETMOMColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGETMOMWith applies the HasEdge predicate on the "GETMOM" edge with a given conditions (other predicates).
func HasGETMOMWith(preds ...predicate.Pregnant) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GETMOMInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GETMOMTable, GETMOMColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTAKECARE applies the HasEdge predicate on the "TAKECARE" edge.
func HasTAKECARE() predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TAKECARETable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TAKECARETable, TAKECAREColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTAKECAREWith applies the HasEdge predicate on the "TAKECARE" edge with a given conditions (other predicates).
func HasTAKECAREWith(preds ...predicate.User) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TAKECAREInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TAKECARETable, TAKECAREColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGETSTATUS applies the HasEdge predicate on the "GETSTATUS" edge.
func HasGETSTATUS() predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GETSTATUSTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GETSTATUSTable, GETSTATUSColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGETSTATUSWith applies the HasEdge predicate on the "GETSTATUS" edge with a given conditions (other predicates).
func HasGETSTATUSWith(preds ...predicate.Babystatus) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GETSTATUSInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GETSTATUSTable, GETSTATUSColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Antenatal) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Antenatal) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
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
func Not(p predicate.Antenatal) predicate.Antenatal {
	return predicate.Antenatal(func(s *sql.Selector) {
		p(s.Not())
	})
}
