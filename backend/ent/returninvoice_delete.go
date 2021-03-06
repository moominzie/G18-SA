// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/returninvoice"
)

// ReturninvoiceDelete is the builder for deleting a Returninvoice entity.
type ReturninvoiceDelete struct {
	config
	hooks      []Hook
	mutation   *ReturninvoiceMutation
	predicates []predicate.Returninvoice
}

// Where adds a new predicate to the delete builder.
func (rd *ReturninvoiceDelete) Where(ps ...predicate.Returninvoice) *ReturninvoiceDelete {
	rd.predicates = append(rd.predicates, ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReturninvoiceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReturninvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReturninvoiceDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReturninvoiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: returninvoice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: returninvoice.FieldID,
			},
		},
	}
	if ps := rd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// ReturninvoiceDeleteOne is the builder for deleting a single Returninvoice entity.
type ReturninvoiceDeleteOne struct {
	rd *ReturninvoiceDelete
}

// Exec executes the deletion query.
func (rdo *ReturninvoiceDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{returninvoice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReturninvoiceDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
