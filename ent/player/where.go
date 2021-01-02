// Code generated by entc, DO NOT EDIT.

package player

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/responserms/server/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// TotalMinutes applies equality check predicate on the "total_minutes" field. It's identical to TotalMinutesEQ.
func TotalMinutes(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalMinutes), v))
	})
}

// SessionStartedAt applies equality check predicate on the "session_started_at" field. It's identical to SessionStartedAtEQ.
func SessionStartedAt(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionStartedAt), v))
	})
}

// SessionEndedAt applies equality check predicate on the "session_ended_at" field. It's identical to SessionEndedAtEQ.
func SessionEndedAt(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionEndedAt), v))
	})
}

// LastSeenAt applies equality check predicate on the "last_seen_at" field. It's identical to LastSeenAtEQ.
func LastSeenAt(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TotalMinutesEQ applies the EQ predicate on the "total_minutes" field.
func TotalMinutesEQ(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalMinutes), v))
	})
}

// TotalMinutesNEQ applies the NEQ predicate on the "total_minutes" field.
func TotalMinutesNEQ(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalMinutes), v))
	})
}

// TotalMinutesIn applies the In predicate on the "total_minutes" field.
func TotalMinutesIn(vs ...int) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTotalMinutes), v...))
	})
}

// TotalMinutesNotIn applies the NotIn predicate on the "total_minutes" field.
func TotalMinutesNotIn(vs ...int) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTotalMinutes), v...))
	})
}

// TotalMinutesGT applies the GT predicate on the "total_minutes" field.
func TotalMinutesGT(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalMinutes), v))
	})
}

// TotalMinutesGTE applies the GTE predicate on the "total_minutes" field.
func TotalMinutesGTE(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalMinutes), v))
	})
}

// TotalMinutesLT applies the LT predicate on the "total_minutes" field.
func TotalMinutesLT(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalMinutes), v))
	})
}

// TotalMinutesLTE applies the LTE predicate on the "total_minutes" field.
func TotalMinutesLTE(v int) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalMinutes), v))
	})
}

// SessionStartedAtEQ applies the EQ predicate on the "session_started_at" field.
func SessionStartedAtEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtNEQ applies the NEQ predicate on the "session_started_at" field.
func SessionStartedAtNEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtIn applies the In predicate on the "session_started_at" field.
func SessionStartedAtIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessionStartedAt), v...))
	})
}

// SessionStartedAtNotIn applies the NotIn predicate on the "session_started_at" field.
func SessionStartedAtNotIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessionStartedAt), v...))
	})
}

// SessionStartedAtGT applies the GT predicate on the "session_started_at" field.
func SessionStartedAtGT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtGTE applies the GTE predicate on the "session_started_at" field.
func SessionStartedAtGTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtLT applies the LT predicate on the "session_started_at" field.
func SessionStartedAtLT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtLTE applies the LTE predicate on the "session_started_at" field.
func SessionStartedAtLTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionStartedAt), v))
	})
}

// SessionStartedAtIsNil applies the IsNil predicate on the "session_started_at" field.
func SessionStartedAtIsNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSessionStartedAt)))
	})
}

// SessionStartedAtNotNil applies the NotNil predicate on the "session_started_at" field.
func SessionStartedAtNotNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSessionStartedAt)))
	})
}

// SessionEndedAtEQ applies the EQ predicate on the "session_ended_at" field.
func SessionEndedAtEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtNEQ applies the NEQ predicate on the "session_ended_at" field.
func SessionEndedAtNEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtIn applies the In predicate on the "session_ended_at" field.
func SessionEndedAtIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessionEndedAt), v...))
	})
}

// SessionEndedAtNotIn applies the NotIn predicate on the "session_ended_at" field.
func SessionEndedAtNotIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessionEndedAt), v...))
	})
}

// SessionEndedAtGT applies the GT predicate on the "session_ended_at" field.
func SessionEndedAtGT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtGTE applies the GTE predicate on the "session_ended_at" field.
func SessionEndedAtGTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtLT applies the LT predicate on the "session_ended_at" field.
func SessionEndedAtLT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtLTE applies the LTE predicate on the "session_ended_at" field.
func SessionEndedAtLTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionEndedAt), v))
	})
}

// SessionEndedAtIsNil applies the IsNil predicate on the "session_ended_at" field.
func SessionEndedAtIsNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSessionEndedAt)))
	})
}

// SessionEndedAtNotNil applies the NotNil predicate on the "session_ended_at" field.
func SessionEndedAtNotNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSessionEndedAt)))
	})
}

// LastSeenAtEQ applies the EQ predicate on the "last_seen_at" field.
func LastSeenAtEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtNEQ applies the NEQ predicate on the "last_seen_at" field.
func LastSeenAtNEQ(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtIn applies the In predicate on the "last_seen_at" field.
func LastSeenAtIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtNotIn applies the NotIn predicate on the "last_seen_at" field.
func LastSeenAtNotIn(vs ...time.Time) predicate.Player {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Player(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtGT applies the GT predicate on the "last_seen_at" field.
func LastSeenAtGT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtGTE applies the GTE predicate on the "last_seen_at" field.
func LastSeenAtGTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLT applies the LT predicate on the "last_seen_at" field.
func LastSeenAtLT(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLTE applies the LTE predicate on the "last_seen_at" field.
func LastSeenAtLTE(v time.Time) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtIsNil applies the IsNil predicate on the "last_seen_at" field.
func LastSeenAtIsNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastSeenAt)))
	})
}

// LastSeenAtNotNil applies the NotNil predicate on the "last_seen_at" field.
func LastSeenAtNotNil() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastSeenAt)))
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasServers applies the HasEdge predicate on the "servers" edge.
func HasServers() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ServersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ServersTable, ServersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServersWith applies the HasEdge predicate on the "servers" edge with a given conditions (other predicates).
func HasServersWith(preds ...predicate.Server) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ServersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ServersTable, ServersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
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

// HasIdentifiers applies the HasEdge predicate on the "identifiers" edge.
func HasIdentifiers() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IdentifiersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IdentifiersTable, IdentifiersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIdentifiersWith applies the HasEdge predicate on the "identifiers" edge with a given conditions (other predicates).
func HasIdentifiersWith(preds ...predicate.PlayerIdentifier) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IdentifiersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IdentifiersTable, IdentifiersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
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
func Not(p predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		p(s.Not())
	})
}
