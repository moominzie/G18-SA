// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/bill"
	"github.com/moominzie/user-record/ent/billingstatus"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetPrice sets the price field.
func (bc *BillCreate) SetPrice(i int) *BillCreate {
	bc.mutation.SetPrice(i)
	return bc
}

// SetTime sets the time field.
func (bc *BillCreate) SetTime(i int) *BillCreate {
	bc.mutation.SetTime(i)
	return bc
}

// SetRepairinvoiceID sets the Repairinvoice edge to RepairInvoice by id.
func (bc *BillCreate) SetRepairinvoiceID(id int) *BillCreate {
	bc.mutation.SetRepairinvoiceID(id)
	return bc
}

// SetNillableRepairinvoiceID sets the Repairinvoice edge to RepairInvoice by id if the given value is not nil.
func (bc *BillCreate) SetNillableRepairinvoiceID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetRepairinvoiceID(*id)
	}
	return bc
}

// SetRepairinvoice sets the Repairinvoice edge to RepairInvoice.
func (bc *BillCreate) SetRepairinvoice(r *RepairInvoice) *BillCreate {
	return bc.SetRepairinvoiceID(r.ID)
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (bc *BillCreate) SetEmployeeID(id int) *BillCreate {
	bc.mutation.SetEmployeeID(id)
	return bc
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (bc *BillCreate) SetNillableEmployeeID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetEmployeeID(*id)
	}
	return bc
}

// SetEmployee sets the Employee edge to Employee.
func (bc *BillCreate) SetEmployee(e *Employee) *BillCreate {
	return bc.SetEmployeeID(e.ID)
}

// SetBillingstatusID sets the Billingstatus edge to Billingstatus by id.
func (bc *BillCreate) SetBillingstatusID(id int) *BillCreate {
	bc.mutation.SetBillingstatusID(id)
	return bc
}

// SetNillableBillingstatusID sets the Billingstatus edge to Billingstatus by id if the given value is not nil.
func (bc *BillCreate) SetNillableBillingstatusID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetBillingstatusID(*id)
	}
	return bc
}

// SetBillingstatus sets the Billingstatus edge to Billingstatus.
func (bc *BillCreate) SetBillingstatus(b *Billingstatus) *BillCreate {
	return bc.SetBillingstatusID(b.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bc *BillCreate) Mutation() *BillMutation {
	return bc.mutation
}

// Save creates the Bill in the database.
func (bc *BillCreate) Save(ctx context.Context) (*Bill, error) {
	if _, ok := bc.mutation.Price(); !ok {
		return nil, &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := bc.mutation.Time(); !ok {
		return nil, &ValidationError{Name: "time", err: errors.New("ent: missing required field \"time\"")}
	}
	var (
		err  error
		node *Bill
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BillCreate) SaveX(ctx context.Context) *Bill {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BillCreate) sqlSave(ctx context.Context) (*Bill, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BillCreate) createSpec() (*Bill, *sqlgraph.CreateSpec) {
	var (
		b     = &Bill{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldPrice,
		})
		b.Price = value
	}
	if value, ok := bc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldTime,
		})
		b.Time = value
	}
	if nodes := bc.mutation.RepairinvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.RepairinvoiceTable,
			Columns: []string{bill.RepairinvoiceColumn},
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
	if nodes := bc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.EmployeeTable,
			Columns: []string{bill.EmployeeColumn},
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
	if nodes := bc.mutation.BillingstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.BillingstatusTable,
			Columns: []string{bill.BillingstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: billingstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}