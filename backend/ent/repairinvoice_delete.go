// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

// RepairInvoiceDelete is the builder for deleting a RepairInvoice entity.
type RepairInvoiceDelete struct {
	config
	hooks      []Hook
	mutation   *RepairInvoiceMutation
	predicates []predicate.RepairInvoice
}

// Where adds a new predicate to the delete builder.
func (rid *RepairInvoiceDelete) Where(ps ...predicate.RepairInvoice) *RepairInvoiceDelete {
	rid.predicates = append(rid.predicates, ps...)
	return rid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rid *RepairInvoiceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rid.hooks) == 0 {
		affected, err = rid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepairInvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rid.mutation = mutation
			affected, err = rid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rid.hooks) - 1; i >= 0; i-- {
			mut = rid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rid *RepairInvoiceDelete) ExecX(ctx context.Context) int {
	n, err := rid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rid *RepairInvoiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: repairinvoice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		},
	}
	if ps := rid.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rid.driver, _spec)
}

// RepairInvoiceDeleteOne is the builder for deleting a single RepairInvoice entity.
type RepairInvoiceDeleteOne struct {
	rid *RepairInvoiceDelete
}

// Exec executes the deletion query.
func (rido *RepairInvoiceDeleteOne) Exec(ctx context.Context) error {
	n, err := rido.rid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{repairinvoice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rido *RepairInvoiceDeleteOne) ExecX(ctx context.Context) {
	rido.rid.ExecX(ctx)
}
