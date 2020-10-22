// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/building"
)

// Building is the model entity for the Building schema.
type Building struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Buname holds the value of the "buname" field.
	Buname string `json:"buname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BuildingQuery when eager-loading is set.
	Edges BuildingEdges `json:"edges"`
}

// BuildingEdges holds the relations/edges for other nodes in the graph.
type BuildingEdges struct {
	// Rooms holds the value of the rooms edge.
	Rooms []*Room
	// UserInformations holds the value of the user_informations edge.
	UserInformations []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading.
func (e BuildingEdges) RoomsOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// UserInformationsOrErr returns the UserInformations value or an error if the edge
// was not loaded in eager-loading.
func (e BuildingEdges) UserInformationsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.UserInformations, nil
	}
	return nil, &NotLoadedError{edge: "user_informations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Building) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // buname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Building fields.
func (b *Building) assignValues(values ...interface{}) error {
	if m, n := len(values), len(building.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field buname", values[0])
	} else if value.Valid {
		b.Buname = value.String
	}
	return nil
}

// QueryRooms queries the rooms edge of the Building.
func (b *Building) QueryRooms() *RoomQuery {
	return (&BuildingClient{config: b.config}).QueryRooms(b)
}

// QueryUserInformations queries the user_informations edge of the Building.
func (b *Building) QueryUserInformations() *UserQuery {
	return (&BuildingClient{config: b.config}).QueryUserInformations(b)
}

// Update returns a builder for updating this Building.
// Note that, you need to call Building.Unwrap() before calling this method, if this Building
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Building) Update() *BuildingUpdateOne {
	return (&BuildingClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Building) Unwrap() *Building {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Building is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Building) String() string {
	var builder strings.Builder
	builder.WriteString("Building(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", buname=")
	builder.WriteString(b.Buname)
	builder.WriteByte(')')
	return builder.String()
}

// Buildings is a parsable slice of Building.
type Buildings []*Building

func (b Buildings) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
