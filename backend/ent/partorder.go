// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/part"
	"github.com/moominzie/user-record/ent/partorder"
	"github.com/moominzie/user-record/ent/repairinvoice"
)

// Partorder is the model entity for the Partorder schema.
type Partorder struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartorderQuery when eager-loading is set.
	Edges           PartorderEdges `json:"edges"`
	employee_id     *int
	part_id         *int
	reparinvoice_id *int
}

// PartorderEdges holds the relations/edges for other nodes in the graph.
type PartorderEdges struct {
	// Repairinvoice holds the value of the repairinvoice edge.
	Repairinvoice *RepairInvoice
	// Employee holds the value of the employee edge.
	Employee *Employee
	// Part holds the value of the part edge.
	Part *Part
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RepairinvoiceOrErr returns the Repairinvoice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartorderEdges) RepairinvoiceOrErr() (*RepairInvoice, error) {
	if e.loadedTypes[0] {
		if e.Repairinvoice == nil {
			// The edge repairinvoice was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repairinvoice.Label}
		}
		return e.Repairinvoice, nil
	}
	return nil, &NotLoadedError{edge: "repairinvoice"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartorderEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[1] {
		if e.Employee == nil {
			// The edge employee was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// PartOrErr returns the Part value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartorderEdges) PartOrErr() (*Part, error) {
	if e.loadedTypes[2] {
		if e.Part == nil {
			// The edge part was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: part.Label}
		}
		return e.Part, nil
	}
	return nil, &NotLoadedError{edge: "part"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Partorder) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Partorder) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // employee_id
		&sql.NullInt64{}, // part_id
		&sql.NullInt64{}, // reparinvoice_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Partorder fields.
func (pa *Partorder) assignValues(values ...interface{}) error {
	if m, n := len(values), len(partorder.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	values = values[0:]
	if len(values) == len(partorder.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field employee_id", value)
		} else if value.Valid {
			pa.employee_id = new(int)
			*pa.employee_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field part_id", value)
		} else if value.Valid {
			pa.part_id = new(int)
			*pa.part_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field reparinvoice_id", value)
		} else if value.Valid {
			pa.reparinvoice_id = new(int)
			*pa.reparinvoice_id = int(value.Int64)
		}
	}
	return nil
}

// QueryRepairinvoice queries the repairinvoice edge of the Partorder.
func (pa *Partorder) QueryRepairinvoice() *RepairInvoiceQuery {
	return (&PartorderClient{config: pa.config}).QueryRepairinvoice(pa)
}

// QueryEmployee queries the employee edge of the Partorder.
func (pa *Partorder) QueryEmployee() *EmployeeQuery {
	return (&PartorderClient{config: pa.config}).QueryEmployee(pa)
}

// QueryPart queries the part edge of the Partorder.
func (pa *Partorder) QueryPart() *PartQuery {
	return (&PartorderClient{config: pa.config}).QueryPart(pa)
}

// Update returns a builder for updating this Partorder.
// Note that, you need to call Partorder.Unwrap() before calling this method, if this Partorder
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Partorder) Update() *PartorderUpdateOne {
	return (&PartorderClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Partorder) Unwrap() *Partorder {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Partorder is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Partorder) String() string {
	var builder strings.Builder
	builder.WriteString("Partorder(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Partorders is a parsable slice of Partorder.
type Partorders []*Partorder

func (pa Partorders) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
