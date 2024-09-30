// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kenta1234567893/upsider-coding-test/ent/company"
	"github.com/kenta1234567893/upsider-coding-test/ent/invoice"
	"github.com/kenta1234567893/upsider-coding-test/ent/predicate"
	"github.com/kenta1234567893/upsider-coding-test/ent/user"
)

// CompanyUpdate is the builder for updating Company entities.
type CompanyUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyMutation
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cu *CompanyUpdate) Where(ps ...predicate.Company) *CompanyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CompanyUpdate) SetUpdatedAt(t time.Time) *CompanyUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *CompanyUpdate) ClearUpdatedAt() *CompanyUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// SetCompanyName sets the "company_name" field.
func (cu *CompanyUpdate) SetCompanyName(s string) *CompanyUpdate {
	cu.mutation.SetCompanyName(s)
	return cu
}

// SetNillableCompanyName sets the "company_name" field if the given value is not nil.
func (cu *CompanyUpdate) SetNillableCompanyName(s *string) *CompanyUpdate {
	if s != nil {
		cu.SetCompanyName(*s)
	}
	return cu
}

// SetCeoName sets the "ceo_name" field.
func (cu *CompanyUpdate) SetCeoName(s string) *CompanyUpdate {
	cu.mutation.SetCeoName(s)
	return cu
}

// SetNillableCeoName sets the "ceo_name" field if the given value is not nil.
func (cu *CompanyUpdate) SetNillableCeoName(s *string) *CompanyUpdate {
	if s != nil {
		cu.SetCeoName(*s)
	}
	return cu
}

// SetPhoneNumber sets the "phone_number" field.
func (cu *CompanyUpdate) SetPhoneNumber(s string) *CompanyUpdate {
	cu.mutation.SetPhoneNumber(s)
	return cu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (cu *CompanyUpdate) SetNillablePhoneNumber(s *string) *CompanyUpdate {
	if s != nil {
		cu.SetPhoneNumber(*s)
	}
	return cu
}

// SetZipCode sets the "zip_code" field.
func (cu *CompanyUpdate) SetZipCode(s string) *CompanyUpdate {
	cu.mutation.SetZipCode(s)
	return cu
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (cu *CompanyUpdate) SetNillableZipCode(s *string) *CompanyUpdate {
	if s != nil {
		cu.SetZipCode(*s)
	}
	return cu
}

// SetAddress sets the "address" field.
func (cu *CompanyUpdate) SetAddress(s string) *CompanyUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cu *CompanyUpdate) SetNillableAddress(s *string) *CompanyUpdate {
	if s != nil {
		cu.SetAddress(*s)
	}
	return cu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cu *CompanyUpdate) AddUserIDs(ids ...uint64) *CompanyUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUsers adds the "users" edges to the User entity.
func (cu *CompanyUpdate) AddUsers(u ...*User) *CompanyUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (cu *CompanyUpdate) AddInvoiceIDs(ids ...uint64) *CompanyUpdate {
	cu.mutation.AddInvoiceIDs(ids...)
	return cu
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (cu *CompanyUpdate) AddInvoices(i ...*Invoice) *CompanyUpdate {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddInvoiceIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cu *CompanyUpdate) Mutation() *CompanyMutation {
	return cu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cu *CompanyUpdate) ClearUsers() *CompanyUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cu *CompanyUpdate) RemoveUserIDs(ids ...uint64) *CompanyUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUsers removes "users" edges to User entities.
func (cu *CompanyUpdate) RemoveUsers(u ...*User) *CompanyUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// ClearInvoices clears all "invoices" edges to the Invoice entity.
func (cu *CompanyUpdate) ClearInvoices() *CompanyUpdate {
	cu.mutation.ClearInvoices()
	return cu
}

// RemoveInvoiceIDs removes the "invoices" edge to Invoice entities by IDs.
func (cu *CompanyUpdate) RemoveInvoiceIDs(ids ...uint64) *CompanyUpdate {
	cu.mutation.RemoveInvoiceIDs(ids...)
	return cu
}

// RemoveInvoices removes "invoices" edges to Invoice entities.
func (cu *CompanyUpdate) RemoveInvoices(i ...*Invoice) *CompanyUpdate {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveInvoiceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompanyUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompanyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompanyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CompanyUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok && !cu.mutation.UpdatedAtCleared() {
		v := company.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(company.Table, company.Columns, sqlgraph.NewFieldSpec(company.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.CreatedAtCleared() {
		_spec.ClearField(company.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(company.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(company.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.CompanyName(); ok {
		_spec.SetField(company.FieldCompanyName, field.TypeString, value)
	}
	if value, ok := cu.mutation.CeoName(); ok {
		_spec.SetField(company.FieldCeoName, field.TypeString, value)
	}
	if value, ok := cu.mutation.PhoneNumber(); ok {
		_spec.SetField(company.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := cu.mutation.ZipCode(); ok {
		_spec.SetField(company.FieldZipCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.SetField(company.FieldAddress, field.TypeString, value)
	}
	if cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedInvoicesIDs(); len(nodes) > 0 && !cu.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CompanyUpdateOne is the builder for updating a single Company entity.
type CompanyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CompanyUpdateOne) SetUpdatedAt(t time.Time) *CompanyUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *CompanyUpdateOne) ClearUpdatedAt() *CompanyUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// SetCompanyName sets the "company_name" field.
func (cuo *CompanyUpdateOne) SetCompanyName(s string) *CompanyUpdateOne {
	cuo.mutation.SetCompanyName(s)
	return cuo
}

// SetNillableCompanyName sets the "company_name" field if the given value is not nil.
func (cuo *CompanyUpdateOne) SetNillableCompanyName(s *string) *CompanyUpdateOne {
	if s != nil {
		cuo.SetCompanyName(*s)
	}
	return cuo
}

// SetCeoName sets the "ceo_name" field.
func (cuo *CompanyUpdateOne) SetCeoName(s string) *CompanyUpdateOne {
	cuo.mutation.SetCeoName(s)
	return cuo
}

// SetNillableCeoName sets the "ceo_name" field if the given value is not nil.
func (cuo *CompanyUpdateOne) SetNillableCeoName(s *string) *CompanyUpdateOne {
	if s != nil {
		cuo.SetCeoName(*s)
	}
	return cuo
}

// SetPhoneNumber sets the "phone_number" field.
func (cuo *CompanyUpdateOne) SetPhoneNumber(s string) *CompanyUpdateOne {
	cuo.mutation.SetPhoneNumber(s)
	return cuo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (cuo *CompanyUpdateOne) SetNillablePhoneNumber(s *string) *CompanyUpdateOne {
	if s != nil {
		cuo.SetPhoneNumber(*s)
	}
	return cuo
}

// SetZipCode sets the "zip_code" field.
func (cuo *CompanyUpdateOne) SetZipCode(s string) *CompanyUpdateOne {
	cuo.mutation.SetZipCode(s)
	return cuo
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (cuo *CompanyUpdateOne) SetNillableZipCode(s *string) *CompanyUpdateOne {
	if s != nil {
		cuo.SetZipCode(*s)
	}
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CompanyUpdateOne) SetAddress(s string) *CompanyUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cuo *CompanyUpdateOne) SetNillableAddress(s *string) *CompanyUpdateOne {
	if s != nil {
		cuo.SetAddress(*s)
	}
	return cuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cuo *CompanyUpdateOne) AddUserIDs(ids ...uint64) *CompanyUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUsers adds the "users" edges to the User entity.
func (cuo *CompanyUpdateOne) AddUsers(u ...*User) *CompanyUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (cuo *CompanyUpdateOne) AddInvoiceIDs(ids ...uint64) *CompanyUpdateOne {
	cuo.mutation.AddInvoiceIDs(ids...)
	return cuo
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (cuo *CompanyUpdateOne) AddInvoices(i ...*Invoice) *CompanyUpdateOne {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddInvoiceIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cuo *CompanyUpdateOne) Mutation() *CompanyMutation {
	return cuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cuo *CompanyUpdateOne) ClearUsers() *CompanyUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cuo *CompanyUpdateOne) RemoveUserIDs(ids ...uint64) *CompanyUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUsers removes "users" edges to User entities.
func (cuo *CompanyUpdateOne) RemoveUsers(u ...*User) *CompanyUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// ClearInvoices clears all "invoices" edges to the Invoice entity.
func (cuo *CompanyUpdateOne) ClearInvoices() *CompanyUpdateOne {
	cuo.mutation.ClearInvoices()
	return cuo
}

// RemoveInvoiceIDs removes the "invoices" edge to Invoice entities by IDs.
func (cuo *CompanyUpdateOne) RemoveInvoiceIDs(ids ...uint64) *CompanyUpdateOne {
	cuo.mutation.RemoveInvoiceIDs(ids...)
	return cuo
}

// RemoveInvoices removes "invoices" edges to Invoice entities.
func (cuo *CompanyUpdateOne) RemoveInvoices(i ...*Invoice) *CompanyUpdateOne {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveInvoiceIDs(ids...)
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cuo *CompanyUpdateOne) Where(ps ...predicate.Company) *CompanyUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompanyUpdateOne) Select(field string, fields ...string) *CompanyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Company entity.
func (cuo *CompanyUpdateOne) Save(ctx context.Context) (*Company, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompanyUpdateOne) SaveX(ctx context.Context) *Company {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompanyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CompanyUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok && !cuo.mutation.UpdatedAtCleared() {
		v := company.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CompanyUpdateOne) sqlSave(ctx context.Context) (_node *Company, err error) {
	_spec := sqlgraph.NewUpdateSpec(company.Table, company.Columns, sqlgraph.NewFieldSpec(company.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Company.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for _, f := range fields {
			if !company.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != company.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.CreatedAtCleared() {
		_spec.ClearField(company.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(company.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(company.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.CompanyName(); ok {
		_spec.SetField(company.FieldCompanyName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CeoName(); ok {
		_spec.SetField(company.FieldCeoName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PhoneNumber(); ok {
		_spec.SetField(company.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ZipCode(); ok {
		_spec.SetField(company.FieldZipCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.SetField(company.FieldAddress, field.TypeString, value)
	}
	if cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedInvoicesIDs(); len(nodes) > 0 && !cuo.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.InvoicesTable,
			Columns: []string{company.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Company{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
