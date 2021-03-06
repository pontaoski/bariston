// Code generated by entc, DO NOT EDIT.

package warning

import (
	"baritone/ent/predicate"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReason), v))
	})
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.Warning {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warning(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReason), v...))
	})
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.Warning {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warning(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReason), v...))
	})
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReason), v))
	})
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReason), v))
	})
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReason), v))
	})
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReason), v))
	})
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReason), v))
	})
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReason), v))
	})
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReason), v))
	})
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReason), v))
	})
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReason), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Warning {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warning(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Warning {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warning(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIssuedBy applies the HasEdge predicate on the "issuedBy" edge.
func HasIssuedBy() predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IssuedByTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, IssuedByTable, IssuedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssuedByWith applies the HasEdge predicate on the "issuedBy" edge with a given conditions (other predicates).
func HasIssuedByWith(preds ...predicate.User) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IssuedByInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, IssuedByTable, IssuedByColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGuild applies the HasEdge predicate on the "guild" edge.
func HasGuild() predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGuildWith applies the HasEdge predicate on the "guild" edge with a given conditions (other predicates).
func HasGuildWith(preds ...predicate.Guild) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Warning) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Warning) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
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
func Not(p predicate.Warning) predicate.Warning {
	return predicate.Warning(func(s *sql.Selector) {
		p(s.Not())
	})
}
