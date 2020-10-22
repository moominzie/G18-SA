// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/statusr"
)

// StatusRDelete is the builder for deleting a StatusR entity.
type StatusRDelete struct {
	config
	hooks      []Hook
	mutation   *StatusRMutation
	predicates []predicate.StatusR
}

// Where adds a new predicate to the delete builder.
func (sr *StatusRDelete) Where(ps ...predicate.StatusR) *StatusRDelete {
	sr.predicates = append(sr.predicates, ps...)
	return sr
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sr *StatusRDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sr.hooks) == 0 {
		affected, err = sr.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusRMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sr.mutation = mutation
			affected, err = sr.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sr.hooks) - 1; i >= 0; i-- {
			mut = sr.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sr.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sr *StatusRDelete) ExecX(ctx context.Context) int {
	n, err := sr.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sr *StatusRDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: statusr.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusr.FieldID,
			},
		},
	}
	if ps := sr.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sr.driver, _spec)
}

// StatusRDeleteOne is the builder for deleting a single StatusR entity.
type StatusRDeleteOne struct {
	sr *StatusRDelete
}

// Exec executes the deletion query.
func (sro *StatusRDeleteOne) Exec(ctx context.Context) error {
	n, err := sro.sr.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{statusr.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sro *StatusRDeleteOne) ExecX(ctx context.Context) {
	sro.sr.ExecX(ctx)
}
