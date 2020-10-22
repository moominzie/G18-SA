// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/repairinvoice"
	"github.com/moominzie/user-record/ent/returninvoice"
	"github.com/moominzie/user-record/ent/statust"
)

// ReturninvoiceUpdate is the builder for updating Returninvoice entities.
type ReturninvoiceUpdate struct {
	config
	hooks      []Hook
	mutation   *ReturninvoiceMutation
	predicates []predicate.Returninvoice
}

// Where adds a new predicate for the builder.
func (ru *ReturninvoiceUpdate) Where(ps ...predicate.Returninvoice) *ReturninvoiceUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetAddedtime sets the addedtime field.
func (ru *ReturninvoiceUpdate) SetAddedtime(t time.Time) *ReturninvoiceUpdate {
	ru.mutation.SetAddedtime(t)
	return ru
}

// SetRepairinvoiceID sets the Repairinvoice edge to RepairInvoice by id.
func (ru *ReturninvoiceUpdate) SetRepairinvoiceID(id int) *ReturninvoiceUpdate {
	ru.mutation.SetRepairinvoiceID(id)
	return ru
}

// SetRepairinvoice sets the Repairinvoice edge to RepairInvoice.
func (ru *ReturninvoiceUpdate) SetRepairinvoice(r *RepairInvoice) *ReturninvoiceUpdate {
	return ru.SetRepairinvoiceID(r.ID)
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (ru *ReturninvoiceUpdate) SetEmployeeID(id int) *ReturninvoiceUpdate {
	ru.mutation.SetEmployeeID(id)
	return ru
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (ru *ReturninvoiceUpdate) SetNillableEmployeeID(id *int) *ReturninvoiceUpdate {
	if id != nil {
		ru = ru.SetEmployeeID(*id)
	}
	return ru
}

// SetEmployee sets the Employee edge to Employee.
func (ru *ReturninvoiceUpdate) SetEmployee(e *Employee) *ReturninvoiceUpdate {
	return ru.SetEmployeeID(e.ID)
}

// SetStatustID sets the Statust edge to Statust by id.
func (ru *ReturninvoiceUpdate) SetStatustID(id int) *ReturninvoiceUpdate {
	ru.mutation.SetStatustID(id)
	return ru
}

// SetNillableStatustID sets the Statust edge to Statust by id if the given value is not nil.
func (ru *ReturninvoiceUpdate) SetNillableStatustID(id *int) *ReturninvoiceUpdate {
	if id != nil {
		ru = ru.SetStatustID(*id)
	}
	return ru
}

// SetStatust sets the Statust edge to Statust.
func (ru *ReturninvoiceUpdate) SetStatust(s *Statust) *ReturninvoiceUpdate {
	return ru.SetStatustID(s.ID)
}

// Mutation returns the ReturninvoiceMutation object of the builder.
func (ru *ReturninvoiceUpdate) Mutation() *ReturninvoiceMutation {
	return ru.mutation
}

// ClearRepairinvoice clears the Repairinvoice edge to RepairInvoice.
func (ru *ReturninvoiceUpdate) ClearRepairinvoice() *ReturninvoiceUpdate {
	ru.mutation.ClearRepairinvoice()
	return ru
}

// ClearEmployee clears the Employee edge to Employee.
func (ru *ReturninvoiceUpdate) ClearEmployee() *ReturninvoiceUpdate {
	ru.mutation.ClearEmployee()
	return ru
}

// ClearStatust clears the Statust edge to Statust.
func (ru *ReturninvoiceUpdate) ClearStatust() *ReturninvoiceUpdate {
	ru.mutation.ClearStatust()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ReturninvoiceUpdate) Save(ctx context.Context) (int, error) {

	if _, ok := ru.mutation.RepairinvoiceID(); ru.mutation.RepairinvoiceCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"Repairinvoice\"")
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReturninvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReturninvoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReturninvoiceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReturninvoiceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReturninvoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   returninvoice.Table,
			Columns: returninvoice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: returninvoice.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: returninvoice.FieldAddedtime,
		})
	}
	if ru.mutation.RepairinvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   returninvoice.RepairinvoiceTable,
			Columns: []string{returninvoice.RepairinvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RepairinvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   returninvoice.RepairinvoiceTable,
			Columns: []string{returninvoice.RepairinvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.EmployeeTable,
			Columns: []string{returninvoice.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.EmployeeTable,
			Columns: []string{returninvoice.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.StatustCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.StatustTable,
			Columns: []string{returninvoice.StatustColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statust.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StatustIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.StatustTable,
			Columns: []string{returninvoice.StatustColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statust.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{returninvoice.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ReturninvoiceUpdateOne is the builder for updating a single Returninvoice entity.
type ReturninvoiceUpdateOne struct {
	config
	hooks    []Hook
	mutation *ReturninvoiceMutation
}

// SetAddedtime sets the addedtime field.
func (ruo *ReturninvoiceUpdateOne) SetAddedtime(t time.Time) *ReturninvoiceUpdateOne {
	ruo.mutation.SetAddedtime(t)
	return ruo
}

// SetRepairinvoiceID sets the Repairinvoice edge to RepairInvoice by id.
func (ruo *ReturninvoiceUpdateOne) SetRepairinvoiceID(id int) *ReturninvoiceUpdateOne {
	ruo.mutation.SetRepairinvoiceID(id)
	return ruo
}

// SetRepairinvoice sets the Repairinvoice edge to RepairInvoice.
func (ruo *ReturninvoiceUpdateOne) SetRepairinvoice(r *RepairInvoice) *ReturninvoiceUpdateOne {
	return ruo.SetRepairinvoiceID(r.ID)
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (ruo *ReturninvoiceUpdateOne) SetEmployeeID(id int) *ReturninvoiceUpdateOne {
	ruo.mutation.SetEmployeeID(id)
	return ruo
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (ruo *ReturninvoiceUpdateOne) SetNillableEmployeeID(id *int) *ReturninvoiceUpdateOne {
	if id != nil {
		ruo = ruo.SetEmployeeID(*id)
	}
	return ruo
}

// SetEmployee sets the Employee edge to Employee.
func (ruo *ReturninvoiceUpdateOne) SetEmployee(e *Employee) *ReturninvoiceUpdateOne {
	return ruo.SetEmployeeID(e.ID)
}

// SetStatustID sets the Statust edge to Statust by id.
func (ruo *ReturninvoiceUpdateOne) SetStatustID(id int) *ReturninvoiceUpdateOne {
	ruo.mutation.SetStatustID(id)
	return ruo
}

// SetNillableStatustID sets the Statust edge to Statust by id if the given value is not nil.
func (ruo *ReturninvoiceUpdateOne) SetNillableStatustID(id *int) *ReturninvoiceUpdateOne {
	if id != nil {
		ruo = ruo.SetStatustID(*id)
	}
	return ruo
}

// SetStatust sets the Statust edge to Statust.
func (ruo *ReturninvoiceUpdateOne) SetStatust(s *Statust) *ReturninvoiceUpdateOne {
	return ruo.SetStatustID(s.ID)
}

// Mutation returns the ReturninvoiceMutation object of the builder.
func (ruo *ReturninvoiceUpdateOne) Mutation() *ReturninvoiceMutation {
	return ruo.mutation
}

// ClearRepairinvoice clears the Repairinvoice edge to RepairInvoice.
func (ruo *ReturninvoiceUpdateOne) ClearRepairinvoice() *ReturninvoiceUpdateOne {
	ruo.mutation.ClearRepairinvoice()
	return ruo
}

// ClearEmployee clears the Employee edge to Employee.
func (ruo *ReturninvoiceUpdateOne) ClearEmployee() *ReturninvoiceUpdateOne {
	ruo.mutation.ClearEmployee()
	return ruo
}

// ClearStatust clears the Statust edge to Statust.
func (ruo *ReturninvoiceUpdateOne) ClearStatust() *ReturninvoiceUpdateOne {
	ruo.mutation.ClearStatust()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *ReturninvoiceUpdateOne) Save(ctx context.Context) (*Returninvoice, error) {

	if _, ok := ruo.mutation.RepairinvoiceID(); ruo.mutation.RepairinvoiceCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"Repairinvoice\"")
	}

	var (
		err  error
		node *Returninvoice
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReturninvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReturninvoiceUpdateOne) SaveX(ctx context.Context) *Returninvoice {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *ReturninvoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReturninvoiceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReturninvoiceUpdateOne) sqlSave(ctx context.Context) (r *Returninvoice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   returninvoice.Table,
			Columns: returninvoice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: returninvoice.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Returninvoice.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: returninvoice.FieldAddedtime,
		})
	}
	if ruo.mutation.RepairinvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   returninvoice.RepairinvoiceTable,
			Columns: []string{returninvoice.RepairinvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RepairinvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   returninvoice.RepairinvoiceTable,
			Columns: []string{returninvoice.RepairinvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.EmployeeTable,
			Columns: []string{returninvoice.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.EmployeeTable,
			Columns: []string{returninvoice.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.StatustCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.StatustTable,
			Columns: []string{returninvoice.StatustColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statust.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StatustIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   returninvoice.StatustTable,
			Columns: []string{returninvoice.StatustColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statust.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Returninvoice{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{returninvoice.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
