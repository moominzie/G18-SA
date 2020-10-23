// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/bill"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/partorder"
	"github.com/moominzie/user-record/ent/returninvoice"
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
}

// SetEmployeename sets the employeename field.
func (ec *EmployeeCreate) SetEmployeename(s string) *EmployeeCreate {
	ec.mutation.SetEmployeename(s)
	return ec
}

// SetEmployeeemail sets the employeeemail field.
func (ec *EmployeeCreate) SetEmployeeemail(s string) *EmployeeCreate {
	ec.mutation.SetEmployeeemail(s)
	return ec
}

// SetPassword sets the password field.
func (ec *EmployeeCreate) SetPassword(s string) *EmployeeCreate {
	ec.mutation.SetPassword(s)
	return ec
}

// AddEmployeeIDs adds the employees edge to Returninvoice by ids.
func (ec *EmployeeCreate) AddEmployeeIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddEmployeeIDs(ids...)
	return ec
}

// AddEmployees adds the employees edges to Returninvoice.
func (ec *EmployeeCreate) AddEmployees(r ...*Returninvoice) *EmployeeCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddEmployeeIDs(ids...)
}

// AddEmployeebillIDs adds the employeebill edge to Bill by ids.
func (ec *EmployeeCreate) AddEmployeebillIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddEmployeebillIDs(ids...)
	return ec
}

// AddEmployeebill adds the employeebill edges to Bill.
func (ec *EmployeeCreate) AddEmployeebill(b ...*Bill) *EmployeeCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ec.AddEmployeebillIDs(ids...)
}

// AddEmployeepartIDs adds the employeepart edge to Partorder by ids.
func (ec *EmployeeCreate) AddEmployeepartIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddEmployeepartIDs(ids...)
	return ec
}

// AddEmployeepart adds the employeepart edges to Partorder.
func (ec *EmployeeCreate) AddEmployeepart(p ...*Partorder) *EmployeeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddEmployeepartIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	if _, ok := ec.mutation.Employeename(); !ok {
		return nil, &ValidationError{Name: "employeename", err: errors.New("ent: missing required field \"employeename\"")}
	}
	if v, ok := ec.mutation.Employeename(); ok {
		if err := employee.EmployeenameValidator(v); err != nil {
			return nil, &ValidationError{Name: "employeename", err: fmt.Errorf("ent: validator failed for field \"employeename\": %w", err)}
		}
	}
	if _, ok := ec.mutation.Employeeemail(); !ok {
		return nil, &ValidationError{Name: "employeeemail", err: errors.New("ent: missing required field \"employeeemail\"")}
	}
	if v, ok := ec.mutation.Employeeemail(); ok {
		if err := employee.EmployeeemailValidator(v); err != nil {
			return nil, &ValidationError{Name: "employeeemail", err: fmt.Errorf("ent: validator failed for field \"employeeemail\": %w", err)}
		}
	}
	if _, ok := ec.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := ec.mutation.Password(); ok {
		if err := employee.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	var (
		err  error
		node *Employee
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		e     = &Employee{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Employeename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeename,
		})
		e.Employeename = value
	}
	if value, ok := ec.mutation.Employeeemail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeemail,
		})
		e.Employeeemail = value
	}
	if value, ok := ec.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldPassword,
		})
		e.Password = value
	}
	if nodes := ec.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeesTable,
			Columns: []string{employee.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: returninvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmployeebillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeebillTable,
			Columns: []string{employee.EmployeebillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmployeepartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeepartTable,
			Columns: []string{employee.EmployeepartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}
