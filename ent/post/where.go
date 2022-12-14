// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kyh0703/stock-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// PublishAt applies equality check predicate on the "publish_at" field. It's identical to PublishAtEQ.
func PublishAt(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublishAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBody), v))
	})
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBody), v))
	})
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBody), v))
	})
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBody), v))
	})
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBody), v))
	})
}

// PublishAtEQ applies the EQ predicate on the "publish_at" field.
func PublishAtEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublishAt), v))
	})
}

// PublishAtNEQ applies the NEQ predicate on the "publish_at" field.
func PublishAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublishAt), v))
	})
}

// PublishAtIn applies the In predicate on the "publish_at" field.
func PublishAtIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPublishAt), v...))
	})
}

// PublishAtNotIn applies the NotIn predicate on the "publish_at" field.
func PublishAtNotIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPublishAt), v...))
	})
}

// PublishAtGT applies the GT predicate on the "publish_at" field.
func PublishAtGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPublishAt), v))
	})
}

// PublishAtGTE applies the GTE predicate on the "publish_at" field.
func PublishAtGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPublishAt), v))
	})
}

// PublishAtLT applies the LT predicate on the "publish_at" field.
func PublishAtLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPublishAt), v))
	})
}

// PublishAtLTE applies the LTE predicate on the "publish_at" field.
func PublishAtLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPublishAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
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
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}
