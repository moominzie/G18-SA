// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/part"
)

// Part is the model entity for the Part schema.
type Part struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pname holds the value of the "Pname" field.
	Pname string `json:"Pname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartQuery when eager-loading is set.
	Edges PartEdges `json:"edges"`
}

// PartEdges holds the relations/edges for other nodes in the graph.
type PartEdges struct {
	// PartInformations holds the value of the part_informations edge.
	PartInformations []*Partorder
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PartInformationsOrErr returns the PartInformations value or an error if the edge
// was not loaded in eager-loading.
func (e PartEdges) PartInformationsOrErr() ([]*Partorder, error) {
	if e.loadedTypes[0] {
		return e.PartInformations, nil
	}
	return nil, &NotLoadedError{edge: "part_informations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Part) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Pname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Part fields.
func (pa *Part) assignValues(values ...interface{}) error {
	if m, n := len(values), len(part.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Pname", values[0])
	} else if value.Valid {
		pa.Pname = value.String
	}
	return nil
}

// QueryPartInformations queries the part_informations edge of the Part.
func (pa *Part) QueryPartInformations() *PartorderQuery {
	return (&PartClient{config: pa.config}).QueryPartInformations(pa)
}

// Update returns a builder for updating this Part.
// Note that, you need to call Part.Unwrap() before calling this method, if this Part
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Part) Update() *PartUpdateOne {
	return (&PartClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Part) Unwrap() *Part {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Part is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Part) String() string {
	var builder strings.Builder
	builder.WriteString("Part(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", Pname=")
	builder.WriteString(pa.Pname)
	builder.WriteByte(')')
	return builder.String()
}

// Parts is a parsable slice of Part.
type Parts []*Part

func (pa Parts) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
