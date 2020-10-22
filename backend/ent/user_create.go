// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/branch"
	"github.com/moominzie/user-record/ent/building"
	"github.com/moominzie/user-record/ent/faculty"
	"github.com/moominzie/user-record/ent/room"
	"github.com/moominzie/user-record/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetPersonalID sets the personalID field.
func (uc *UserCreate) SetPersonalID(s string) *UserCreate {
	uc.mutation.SetPersonalID(s)
	return uc
}

// SetPersonalName sets the personalName field.
func (uc *UserCreate) SetPersonalName(s string) *UserCreate {
	uc.mutation.SetPersonalName(s)
	return uc
}

// SetFacultyID sets the faculty edge to Faculty by id.
func (uc *UserCreate) SetFacultyID(id int) *UserCreate {
	uc.mutation.SetFacultyID(id)
	return uc
}

// SetNillableFacultyID sets the faculty edge to Faculty by id if the given value is not nil.
func (uc *UserCreate) SetNillableFacultyID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetFacultyID(*id)
	}
	return uc
}

// SetFaculty sets the faculty edge to Faculty.
func (uc *UserCreate) SetFaculty(f *Faculty) *UserCreate {
	return uc.SetFacultyID(f.ID)
}

// SetBranchID sets the branch edge to Branch by id.
func (uc *UserCreate) SetBranchID(id int) *UserCreate {
	uc.mutation.SetBranchID(id)
	return uc
}

// SetNillableBranchID sets the branch edge to Branch by id if the given value is not nil.
func (uc *UserCreate) SetNillableBranchID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetBranchID(*id)
	}
	return uc
}

// SetBranch sets the branch edge to Branch.
func (uc *UserCreate) SetBranch(b *Branch) *UserCreate {
	return uc.SetBranchID(b.ID)
}

// SetBuildingID sets the building edge to Building by id.
func (uc *UserCreate) SetBuildingID(id int) *UserCreate {
	uc.mutation.SetBuildingID(id)
	return uc
}

// SetNillableBuildingID sets the building edge to Building by id if the given value is not nil.
func (uc *UserCreate) SetNillableBuildingID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetBuildingID(*id)
	}
	return uc
}

// SetBuilding sets the building edge to Building.
func (uc *UserCreate) SetBuilding(b *Building) *UserCreate {
	return uc.SetBuildingID(b.ID)
}

// SetRoomID sets the room edge to Room by id.
func (uc *UserCreate) SetRoomID(id int) *UserCreate {
	uc.mutation.SetRoomID(id)
	return uc
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (uc *UserCreate) SetNillableRoomID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetRoomID(*id)
	}
	return uc
}

// SetRoom sets the room edge to Room.
func (uc *UserCreate) SetRoom(r *Room) *UserCreate {
	return uc.SetRoomID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.PersonalID(); !ok {
		return nil, &ValidationError{Name: "personalID", err: errors.New("ent: missing required field \"personalID\"")}
	}
	if _, ok := uc.mutation.PersonalName(); !ok {
		return nil, &ValidationError{Name: "personalName", err: errors.New("ent: missing required field \"personalName\"")}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.PersonalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPersonalID,
		})
		u.PersonalID = value
	}
	if value, ok := uc.mutation.PersonalName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPersonalName,
		})
		u.PersonalName = value
	}
	if nodes := uc.mutation.FacultyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.FacultyTable,
			Columns: []string{user.FacultyColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BranchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.BranchTable,
			Columns: []string{user.BranchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: branch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BuildingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.BuildingTable,
			Columns: []string{user.BuildingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: building.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoomTable,
			Columns: []string{user.RoomColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
