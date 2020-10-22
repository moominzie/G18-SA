// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/device"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Dname holds the value of the "Dname" field.
	Dname string `json:"Dname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges DeviceEdges `json:"edges"`
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// RepairInformation holds the value of the repair_information edge.
	RepairInformation []*RepairInvoice
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RepairInformationOrErr returns the RepairInformation value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) RepairInformationOrErr() ([]*RepairInvoice, error) {
	if e.loadedTypes[0] {
		return e.RepairInformation, nil
	}
	return nil, &NotLoadedError{edge: "repair_information"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Dname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(values ...interface{}) error {
	if m, n := len(values), len(device.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Dname", values[0])
	} else if value.Valid {
		d.Dname = value.String
	}
	return nil
}

// QueryRepairInformation queries the repair_information edge of the Device.
func (d *Device) QueryRepairInformation() *RepairInvoiceQuery {
	return (&DeviceClient{config: d.config}).QueryRepairInformation(d)
}

// Update returns a builder for updating this Device.
// Note that, you need to call Device.Unwrap() before calling this method, if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Dname=")
	builder.WriteString(d.Dname)
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
