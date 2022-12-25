// Code generated by ent, DO NOT EDIT.

package tenant

import (
	"time"
)

const (
	// Label holds the string label denoting the tenant type in the database.
	Label = "tenant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldSortNo holds the string denoting the sort_no field in the database.
	FieldSortNo = "sort_no"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the tenant in the database.
	Table = "sys_tenant"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "tenant_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "sys_users"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_tenant"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "tenant_children"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_tenant"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "tenant_children"
)

// Columns holds all SQL columns for tenant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldUUID,
	FieldPid,
	FieldName,
	FieldAccount,
	FieldStartTime,
	FieldEndTime,
	FieldContact,
	FieldMobile,
	FieldSortNo,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_tenant"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tenant_children",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"tenant_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID string
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime func() time.Time
	// DefaultSortNo holds the default value on creation for the "sort_no" field.
	DefaultSortNo int
)
