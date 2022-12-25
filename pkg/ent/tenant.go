// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/pkg/ent/tenant"
)

// Tenant is the model entity for the Tenant schema.
type Tenant struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// status 1 normal 0 ban | 状态 1 正常 0 禁用
	Status uint8 `json:"status,omitempty"`
	// tenant's UUID | 租户的UUID
	UUID string `json:"uuid,omitempty"`
	// parent id | 父级ID
	Pid *string `json:"pid,omitempty"`
	// tenant's name | 租户的名称
	Name string `json:"name,omitempty"`
	// tenant's account | 租户登录账号
	Account string `json:"account,omitempty"`
	// start_time | 租期的开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// end_time ｜ 租期的结束时间
	EndTime time.Time `json:"end_time,omitempty"`
	// contact | 客户联系人
	Contact string `json:"contact,omitempty"`
	// mobile | 客户联系电话
	Mobile string `json:"mobile,omitempty"`
	// sort number | 显示排序
	SortNo int `json:"sort_no,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TenantQuery when eager-loading is set.
	Edges           TenantEdges `json:"edges"`
	tenant_children *uint64
}

// TenantEdges holds the relations/edges for other nodes in the graph.
type TenantEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Children holds the value of the children edge.
	Children []*Tenant `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Tenant `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e TenantEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TenantEdges) ChildrenOrErr() ([]*Tenant, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TenantEdges) ParentOrErr() (*Tenant, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tenant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tenant.FieldID, tenant.FieldStatus, tenant.FieldSortNo:
			values[i] = new(sql.NullInt64)
		case tenant.FieldUUID, tenant.FieldPid, tenant.FieldName, tenant.FieldAccount, tenant.FieldContact, tenant.FieldMobile:
			values[i] = new(sql.NullString)
		case tenant.FieldCreatedAt, tenant.FieldUpdatedAt, tenant.FieldStartTime, tenant.FieldEndTime:
			values[i] = new(sql.NullTime)
		case tenant.ForeignKeys[0]: // tenant_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tenant", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tenant fields.
func (t *Tenant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tenant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = uint64(value.Int64)
		case tenant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tenant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case tenant.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = uint8(value.Int64)
			}
		case tenant.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				t.UUID = value.String
			}
		case tenant.FieldPid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				t.Pid = new(string)
				*t.Pid = value.String
			}
		case tenant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tenant.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				t.Account = value.String
			}
		case tenant.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				t.StartTime = value.Time
			}
		case tenant.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				t.EndTime = value.Time
			}
		case tenant.FieldContact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact", values[i])
			} else if value.Valid {
				t.Contact = value.String
			}
		case tenant.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				t.Mobile = value.String
			}
		case tenant.FieldSortNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_no", values[i])
			} else if value.Valid {
				t.SortNo = int(value.Int64)
			}
		case tenant.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tenant_children", value)
			} else if value.Valid {
				t.tenant_children = new(uint64)
				*t.tenant_children = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Tenant entity.
func (t *Tenant) QueryUsers() *UserQuery {
	return (&TenantClient{config: t.config}).QueryUsers(t)
}

// QueryChildren queries the "children" edge of the Tenant entity.
func (t *Tenant) QueryChildren() *TenantQuery {
	return (&TenantClient{config: t.config}).QueryChildren(t)
}

// QueryParent queries the "parent" edge of the Tenant entity.
func (t *Tenant) QueryParent() *TenantQuery {
	return (&TenantClient{config: t.config}).QueryParent(t)
}

// Update returns a builder for updating this Tenant.
// Note that you need to call Tenant.Unwrap() before calling this method if this Tenant
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tenant) Update() *TenantUpdateOne {
	return (&TenantClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tenant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tenant) Unwrap() *Tenant {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tenant is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tenant) String() string {
	var builder strings.Builder
	builder.WriteString("Tenant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(t.UUID)
	builder.WriteString(", ")
	if v := t.Pid; v != nil {
		builder.WriteString("pid=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("account=")
	builder.WriteString(t.Account)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(t.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(t.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("contact=")
	builder.WriteString(t.Contact)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(t.Mobile)
	builder.WriteString(", ")
	builder.WriteString("sort_no=")
	builder.WriteString(fmt.Sprintf("%v", t.SortNo))
	builder.WriteByte(')')
	return builder.String()
}

// Tenants is a parsable slice of Tenant.
type Tenants []*Tenant

func (t Tenants) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
