// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/building"
	"github.com/moominzie/user-record/ent/predicate"
	"github.com/moominzie/user-record/ent/room"
	"github.com/moominzie/user-record/ent/user"
)

// BuildingUpdate is the builder for updating Building entities.
type BuildingUpdate struct {
	config
	hooks      []Hook
	mutation   *BuildingMutation
	predicates []predicate.Building
}

// Where adds a new predicate for the builder.
func (bu *BuildingUpdate) Where(ps ...predicate.Building) *BuildingUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBuname sets the buname field.
func (bu *BuildingUpdate) SetBuname(s string) *BuildingUpdate {
	bu.mutation.SetBuname(s)
	return bu
}

// AddRoomIDs adds the rooms edge to Room by ids.
func (bu *BuildingUpdate) AddRoomIDs(ids ...int) *BuildingUpdate {
	bu.mutation.AddRoomIDs(ids...)
	return bu
}

// AddRooms adds the rooms edges to Room.
func (bu *BuildingUpdate) AddRooms(r ...*Room) *BuildingUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bu.AddRoomIDs(ids...)
}

// AddUserInformationIDs adds the user_informations edge to User by ids.
func (bu *BuildingUpdate) AddUserInformationIDs(ids ...int) *BuildingUpdate {
	bu.mutation.AddUserInformationIDs(ids...)
	return bu
}

// AddUserInformations adds the user_informations edges to User.
func (bu *BuildingUpdate) AddUserInformations(u ...*User) *BuildingUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddUserInformationIDs(ids...)
}

// Mutation returns the BuildingMutation object of the builder.
func (bu *BuildingUpdate) Mutation() *BuildingMutation {
	return bu.mutation
}

// RemoveRoomIDs removes the rooms edge to Room by ids.
func (bu *BuildingUpdate) RemoveRoomIDs(ids ...int) *BuildingUpdate {
	bu.mutation.RemoveRoomIDs(ids...)
	return bu
}

// RemoveRooms removes rooms edges to Room.
func (bu *BuildingUpdate) RemoveRooms(r ...*Room) *BuildingUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bu.RemoveRoomIDs(ids...)
}

// RemoveUserInformationIDs removes the user_informations edge to User by ids.
func (bu *BuildingUpdate) RemoveUserInformationIDs(ids ...int) *BuildingUpdate {
	bu.mutation.RemoveUserInformationIDs(ids...)
	return bu
}

// RemoveUserInformations removes user_informations edges to User.
func (bu *BuildingUpdate) RemoveUserInformations(u ...*User) *BuildingUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveUserInformationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BuildingUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildingMutation)
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
func (bu *BuildingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BuildingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BuildingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BuildingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   building.Table,
			Columns: building.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: building.FieldID,
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
	if value, ok := bu.mutation.Buname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: building.FieldBuname,
		})
	}
	if nodes := bu.mutation.RemovedRoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   building.RoomsTable,
			Columns: []string{building.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   building.RoomsTable,
			Columns: []string{building.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
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
			Table:   building.UserInformationsTable,
			Columns: []string{building.UserInformationsColumn},
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
			Table:   building.UserInformationsTable,
			Columns: []string{building.UserInformationsColumn},
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
			err = &NotFoundError{building.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BuildingUpdateOne is the builder for updating a single Building entity.
type BuildingUpdateOne struct {
	config
	hooks    []Hook
	mutation *BuildingMutation
}

// SetBuname sets the buname field.
func (buo *BuildingUpdateOne) SetBuname(s string) *BuildingUpdateOne {
	buo.mutation.SetBuname(s)
	return buo
}

// AddRoomIDs adds the rooms edge to Room by ids.
func (buo *BuildingUpdateOne) AddRoomIDs(ids ...int) *BuildingUpdateOne {
	buo.mutation.AddRoomIDs(ids...)
	return buo
}

// AddRooms adds the rooms edges to Room.
func (buo *BuildingUpdateOne) AddRooms(r ...*Room) *BuildingUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return buo.AddRoomIDs(ids...)
}

// AddUserInformationIDs adds the user_informations edge to User by ids.
func (buo *BuildingUpdateOne) AddUserInformationIDs(ids ...int) *BuildingUpdateOne {
	buo.mutation.AddUserInformationIDs(ids...)
	return buo
}

// AddUserInformations adds the user_informations edges to User.
func (buo *BuildingUpdateOne) AddUserInformations(u ...*User) *BuildingUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddUserInformationIDs(ids...)
}

// Mutation returns the BuildingMutation object of the builder.
func (buo *BuildingUpdateOne) Mutation() *BuildingMutation {
	return buo.mutation
}

// RemoveRoomIDs removes the rooms edge to Room by ids.
func (buo *BuildingUpdateOne) RemoveRoomIDs(ids ...int) *BuildingUpdateOne {
	buo.mutation.RemoveRoomIDs(ids...)
	return buo
}

// RemoveRooms removes rooms edges to Room.
func (buo *BuildingUpdateOne) RemoveRooms(r ...*Room) *BuildingUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return buo.RemoveRoomIDs(ids...)
}

// RemoveUserInformationIDs removes the user_informations edge to User by ids.
func (buo *BuildingUpdateOne) RemoveUserInformationIDs(ids ...int) *BuildingUpdateOne {
	buo.mutation.RemoveUserInformationIDs(ids...)
	return buo
}

// RemoveUserInformations removes user_informations edges to User.
func (buo *BuildingUpdateOne) RemoveUserInformations(u ...*User) *BuildingUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveUserInformationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BuildingUpdateOne) Save(ctx context.Context) (*Building, error) {

	var (
		err  error
		node *Building
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildingMutation)
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
func (buo *BuildingUpdateOne) SaveX(ctx context.Context) *Building {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BuildingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BuildingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BuildingUpdateOne) sqlSave(ctx context.Context) (b *Building, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   building.Table,
			Columns: building.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: building.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Building.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Buname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: building.FieldBuname,
		})
	}
	if nodes := buo.mutation.RemovedRoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   building.RoomsTable,
			Columns: []string{building.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   building.RoomsTable,
			Columns: []string{building.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
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
			Table:   building.UserInformationsTable,
			Columns: []string{building.UserInformationsColumn},
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
			Table:   building.UserInformationsTable,
			Columns: []string{building.UserInformationsColumn},
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
	b = &Building{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{building.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}