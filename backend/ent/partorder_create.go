// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/part"
	"github.com/moominzie/user-record/ent/partorder"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

// PartorderCreate is the builder for creating a Partorder entity.
type PartorderCreate struct {
	config
	mutation *PartorderMutation
	hooks    []Hook
}

// SetRepairinvoiceID sets the repairinvoice edge to RepairInvoice by id.
func (pc *PartorderCreate) SetRepairinvoiceID(id int) *PartorderCreate {
	pc.mutation.SetRepairinvoiceID(id)
	return pc
}

// SetNillableRepairinvoiceID sets the repairinvoice edge to RepairInvoice by id if the given value is not nil.
func (pc *PartorderCreate) SetNillableRepairinvoiceID(id *int) *PartorderCreate {
	if id != nil {
		pc = pc.SetRepairinvoiceID(*id)
	}
	return pc
}

// SetRepairinvoice sets the repairinvoice edge to RepairInvoice.
func (pc *PartorderCreate) SetRepairinvoice(r *RepairInvoice) *PartorderCreate {
	return pc.SetRepairinvoiceID(r.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (pc *PartorderCreate) SetEmployeeID(id int) *PartorderCreate {
	pc.mutation.SetEmployeeID(id)
	return pc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (pc *PartorderCreate) SetNillableEmployeeID(id *int) *PartorderCreate {
	if id != nil {
		pc = pc.SetEmployeeID(*id)
	}
	return pc
}

// SetEmployee sets the employee edge to Employee.
func (pc *PartorderCreate) SetEmployee(e *Employee) *PartorderCreate {
	return pc.SetEmployeeID(e.ID)
}

// SetPartID sets the part edge to Part by id.
func (pc *PartorderCreate) SetPartID(id int) *PartorderCreate {
	pc.mutation.SetPartID(id)
	return pc
}

// SetNillablePartID sets the part edge to Part by id if the given value is not nil.
func (pc *PartorderCreate) SetNillablePartID(id *int) *PartorderCreate {
	if id != nil {
		pc = pc.SetPartID(*id)
	}
	return pc
}

// SetPart sets the part edge to Part.
func (pc *PartorderCreate) SetPart(p *Part) *PartorderCreate {
	return pc.SetPartID(p.ID)
}

// Mutation returns the PartorderMutation object of the builder.
func (pc *PartorderCreate) Mutation() *PartorderMutation {
	return pc.mutation
}

// Save creates the Partorder in the database.
func (pc *PartorderCreate) Save(ctx context.Context) (*Partorder, error) {
	var (
		err  error
		node *Partorder
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartorderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PartorderCreate) SaveX(ctx context.Context) *Partorder {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PartorderCreate) sqlSave(ctx context.Context) (*Partorder, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PartorderCreate) createSpec() (*Partorder, *sqlgraph.CreateSpec) {
	var (
		pa    = &Partorder{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: partorder.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partorder.FieldID,
			},
		}
	)
	if nodes := pc.mutation.RepairinvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   partorder.RepairinvoiceTable,
			Columns: []string{partorder.RepairinvoiceColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partorder.EmployeeTable,
			Columns: []string{partorder.EmployeeColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partorder.PartTable,
			Columns: []string{partorder.PartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
