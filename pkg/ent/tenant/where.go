// Code generated by ent, DO NOT EDIT.

package tenant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPid), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Account applies equality check predicate on the "account" field. It's identical to AccountEQ.
func Account(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccount), v))
	})
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// Contact applies equality check predicate on the "contact" field. It's identical to ContactEQ.
func Contact(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContact), v))
	})
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// SortNo applies equality check predicate on the "sort_no" field. It's identical to SortNoEQ.
func SortNo(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSortNo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPid), v))
	})
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPid), v))
	})
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPid), v...))
	})
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPid), v...))
	})
}

// PidGT applies the GT predicate on the "pid" field.
func PidGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPid), v))
	})
}

// PidGTE applies the GTE predicate on the "pid" field.
func PidGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPid), v))
	})
}

// PidLT applies the LT predicate on the "pid" field.
func PidLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPid), v))
	})
}

// PidLTE applies the LTE predicate on the "pid" field.
func PidLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPid), v))
	})
}

// PidContains applies the Contains predicate on the "pid" field.
func PidContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPid), v))
	})
}

// PidHasPrefix applies the HasPrefix predicate on the "pid" field.
func PidHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPid), v))
	})
}

// PidHasSuffix applies the HasSuffix predicate on the "pid" field.
func PidHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPid), v))
	})
}

// PidIsNil applies the IsNil predicate on the "pid" field.
func PidIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPid)))
	})
}

// PidNotNil applies the NotNil predicate on the "pid" field.
func PidNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPid)))
	})
}

// PidEqualFold applies the EqualFold predicate on the "pid" field.
func PidEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPid), v))
	})
}

// PidContainsFold applies the ContainsFold predicate on the "pid" field.
func PidContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPid), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AccountEQ applies the EQ predicate on the "account" field.
func AccountEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccount), v))
	})
}

// AccountNEQ applies the NEQ predicate on the "account" field.
func AccountNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccount), v))
	})
}

// AccountIn applies the In predicate on the "account" field.
func AccountIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccount), v...))
	})
}

// AccountNotIn applies the NotIn predicate on the "account" field.
func AccountNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccount), v...))
	})
}

// AccountGT applies the GT predicate on the "account" field.
func AccountGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccount), v))
	})
}

// AccountGTE applies the GTE predicate on the "account" field.
func AccountGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccount), v))
	})
}

// AccountLT applies the LT predicate on the "account" field.
func AccountLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccount), v))
	})
}

// AccountLTE applies the LTE predicate on the "account" field.
func AccountLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccount), v))
	})
}

// AccountContains applies the Contains predicate on the "account" field.
func AccountContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccount), v))
	})
}

// AccountHasPrefix applies the HasPrefix predicate on the "account" field.
func AccountHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccount), v))
	})
}

// AccountHasSuffix applies the HasSuffix predicate on the "account" field.
func AccountHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccount), v))
	})
}

// AccountEqualFold applies the EqualFold predicate on the "account" field.
func AccountEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccount), v))
	})
}

// AccountContainsFold applies the ContainsFold predicate on the "account" field.
func AccountContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccount), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// EndTimeIsNil applies the IsNil predicate on the "end_time" field.
func EndTimeIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndTime)))
	})
}

// EndTimeNotNil applies the NotNil predicate on the "end_time" field.
func EndTimeNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndTime)))
	})
}

// ContactEQ applies the EQ predicate on the "contact" field.
func ContactEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContact), v))
	})
}

// ContactNEQ applies the NEQ predicate on the "contact" field.
func ContactNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContact), v))
	})
}

// ContactIn applies the In predicate on the "contact" field.
func ContactIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContact), v...))
	})
}

// ContactNotIn applies the NotIn predicate on the "contact" field.
func ContactNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContact), v...))
	})
}

// ContactGT applies the GT predicate on the "contact" field.
func ContactGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContact), v))
	})
}

// ContactGTE applies the GTE predicate on the "contact" field.
func ContactGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContact), v))
	})
}

// ContactLT applies the LT predicate on the "contact" field.
func ContactLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContact), v))
	})
}

// ContactLTE applies the LTE predicate on the "contact" field.
func ContactLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContact), v))
	})
}

// ContactContains applies the Contains predicate on the "contact" field.
func ContactContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContact), v))
	})
}

// ContactHasPrefix applies the HasPrefix predicate on the "contact" field.
func ContactHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContact), v))
	})
}

// ContactHasSuffix applies the HasSuffix predicate on the "contact" field.
func ContactHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContact), v))
	})
}

// ContactIsNil applies the IsNil predicate on the "contact" field.
func ContactIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContact)))
	})
}

// ContactNotNil applies the NotNil predicate on the "contact" field.
func ContactNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContact)))
	})
}

// ContactEqualFold applies the EqualFold predicate on the "contact" field.
func ContactEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContact), v))
	})
}

// ContactContainsFold applies the ContainsFold predicate on the "contact" field.
func ContactContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContact), v))
	})
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMobile), v))
	})
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMobile), v...))
	})
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMobile), v...))
	})
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMobile), v))
	})
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMobile), v))
	})
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMobile), v))
	})
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMobile), v))
	})
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMobile), v))
	})
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMobile), v))
	})
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMobile), v))
	})
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMobile)))
	})
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMobile)))
	})
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMobile), v))
	})
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMobile), v))
	})
}

// SortNoEQ applies the EQ predicate on the "sort_no" field.
func SortNoEQ(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSortNo), v))
	})
}

// SortNoNEQ applies the NEQ predicate on the "sort_no" field.
func SortNoNEQ(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSortNo), v))
	})
}

// SortNoIn applies the In predicate on the "sort_no" field.
func SortNoIn(vs ...int) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSortNo), v...))
	})
}

// SortNoNotIn applies the NotIn predicate on the "sort_no" field.
func SortNoNotIn(vs ...int) predicate.Tenant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSortNo), v...))
	})
}

// SortNoGT applies the GT predicate on the "sort_no" field.
func SortNoGT(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSortNo), v))
	})
}

// SortNoGTE applies the GTE predicate on the "sort_no" field.
func SortNoGTE(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSortNo), v))
	})
}

// SortNoLT applies the LT predicate on the "sort_no" field.
func SortNoLT(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSortNo), v))
	})
}

// SortNoLTE applies the LTE predicate on the "sort_no" field.
func SortNoLTE(v int) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSortNo), v))
	})
}

// SortNoIsNil applies the IsNil predicate on the "sort_no" field.
func SortNoIsNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSortNo)))
	})
}

// SortNoNotNil applies the NotNil predicate on the "sort_no" field.
func SortNoNotNil() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSortNo)))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
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
func Not(p predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		p(s.Not())
	})
}
