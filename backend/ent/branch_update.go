// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/branch"
	"github.com/moominzie/user-record/ent/faculty"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/user"
)

// BranchUpdate is the builder for updating Branch entities.
type BranchUpdate struct {
	config
	hooks      []Hook
	mutation   *BranchMutation
	predicates []predicate.Branch
}

// Where adds a new predicate for the builder.
func (bu *BranchUpdate) Where(ps ...predicate.Branch) *BranchUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBrname sets the brname field.
func (bu *BranchUpdate) SetBrname(s string) *BranchUpdate {
	bu.mutation.SetBrname(s)
	return bu
}

// SetFacultyID sets the faculty edge to Faculty by id.
func (bu *BranchUpdate) SetFacultyID(id int) *BranchUpdate {
	bu.mutation.SetFacultyID(id)
	return bu
}

// SetNillableFacultyID sets the faculty edge to Faculty by id if the given value is not nil.
func (bu *BranchUpdate) SetNillableFacultyID(id *int) *BranchUpdate {
	if id != nil {
		bu = bu.SetFacultyID(*id)
	}
	return bu
}

// SetFaculty sets the faculty edge to Faculty.
func (bu *BranchUpdate) SetFaculty(f *Faculty) *BranchUpdate {
	return bu.SetFacultyID(f.ID)
}

// AddUserInformationIDs adds the user_informations edge to User by ids.
func (bu *BranchUpdate) AddUserInformationIDs(ids ...int) *BranchUpdate {
	bu.mutation.AddUserInformationIDs(ids...)
	return bu
}

// AddUserInformations adds the user_informations edges to User.
func (bu *BranchUpdate) AddUserInformations(u ...*User) *BranchUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddUserInformationIDs(ids...)
}

// Mutation returns the BranchMutation object of the builder.
func (bu *BranchUpdate) Mutation() *BranchMutation {
	return bu.mutation
}

// ClearFaculty clears the faculty edge to Faculty.
func (bu *BranchUpdate) ClearFaculty() *BranchUpdate {
	bu.mutation.ClearFaculty()
	return bu
}

// RemoveUserInformationIDs removes the user_informations edge to User by ids.
func (bu *BranchUpdate) RemoveUserInformationIDs(ids ...int) *BranchUpdate {
	bu.mutation.RemoveUserInformationIDs(ids...)
	return bu
}

// RemoveUserInformations removes user_informations edges to User.
func (bu *BranchUpdate) RemoveUserInformations(u ...*User) *BranchUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveUserInformationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BranchUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BranchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BranchUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BranchUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BranchUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BranchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   branch.Table,
			Columns: branch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branch.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Brname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldBrname,
		})
	}
	if bu.mutation.FacultyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   branch.FacultyTable,
			Columns: []string{branch.FacultyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.FacultyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   branch.FacultyTable,
			Columns: []string{branch.FacultyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := bu.mutation.RemovedUserInformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   branch.UserInformationsTable,
			Columns: []string{branch.UserInformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UserInformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   branch.UserInformationsTable,
			Columns: []string{branch.UserInformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{branch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BranchUpdateOne is the builder for updating a single Branch entity.
type BranchUpdateOne struct {
	config
	hooks    []Hook
	mutation *BranchMutation
}

// SetBrname sets the brname field.
func (buo *BranchUpdateOne) SetBrname(s string) *BranchUpdateOne {
	buo.mutation.SetBrname(s)
	return buo
}

// SetFacultyID sets the faculty edge to Faculty by id.
func (buo *BranchUpdateOne) SetFacultyID(id int) *BranchUpdateOne {
	buo.mutation.SetFacultyID(id)
	return buo
}

// SetNillableFacultyID sets the faculty edge to Faculty by id if the given value is not nil.
func (buo *BranchUpdateOne) SetNillableFacultyID(id *int) *BranchUpdateOne {
	if id != nil {
		buo = buo.SetFacultyID(*id)
	}
	return buo
}

// SetFaculty sets the faculty edge to Faculty.
func (buo *BranchUpdateOne) SetFaculty(f *Faculty) *BranchUpdateOne {
	return buo.SetFacultyID(f.ID)
}

// AddUserInformationIDs adds the user_informations edge to User by ids.
func (buo *BranchUpdateOne) AddUserInformationIDs(ids ...int) *BranchUpdateOne {
	buo.mutation.AddUserInformationIDs(ids...)
	return buo
}

// AddUserInformations adds the user_informations edges to User.
func (buo *BranchUpdateOne) AddUserInformations(u ...*User) *BranchUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddUserInformationIDs(ids...)
}

// Mutation returns the BranchMutation object of the builder.
func (buo *BranchUpdateOne) Mutation() *BranchMutation {
	return buo.mutation
}

// ClearFaculty clears the faculty edge to Faculty.
func (buo *BranchUpdateOne) ClearFaculty() *BranchUpdateOne {
	buo.mutation.ClearFaculty()
	return buo
}

// RemoveUserInformationIDs removes the user_informations edge to User by ids.
func (buo *BranchUpdateOne) RemoveUserInformationIDs(ids ...int) *BranchUpdateOne {
	buo.mutation.RemoveUserInformationIDs(ids...)
	return buo
}

// RemoveUserInformations removes user_informations edges to User.
func (buo *BranchUpdateOne) RemoveUserInformations(u ...*User) *BranchUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveUserInformationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BranchUpdateOne) Save(ctx context.Context) (*Branch, error) {

	var (
		err  error
		node *Branch
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BranchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BranchUpdateOne) SaveX(ctx context.Context) *Branch {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BranchUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BranchUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BranchUpdateOne) sqlSave(ctx context.Context) (b *Branch, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   branch.Table,
			Columns: branch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branch.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Branch.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Brname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldBrname,
		})
	}
	if buo.mutation.FacultyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   branch.FacultyTable,
			Columns: []string{branch.FacultyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.FacultyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   branch.FacultyTable,
			Columns: []string{branch.FacultyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := buo.mutation.RemovedUserInformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   branch.UserInformationsTable,
			Columns: []string{branch.UserInformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UserInformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   branch.UserInformationsTable,
			Columns: []string{branch.UserInformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Branch{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{branch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
