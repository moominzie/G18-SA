// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/branch"
	"github.com/moominzie/user-record/ent/faculty"
)

// Branch is the model entity for the Branch schema.
type Branch struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Brname holds the value of the "brname" field.
	Brname string `json:"brname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BranchQuery when eager-loading is set.
	Edges      BranchEdges `json:"edges"`
	faculty_id *int
}

// BranchEdges holds the relations/edges for other nodes in the graph.
type BranchEdges struct {
	// Faculty holds the value of the faculty edge.
	Faculty *Faculty
	// UserInformations holds the value of the user_informations edge.
	UserInformations []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FacultyOrErr returns the Faculty value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BranchEdges) FacultyOrErr() (*Faculty, error) {
	if e.loadedTypes[0] {
		if e.Faculty == nil {
			// The edge faculty was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: faculty.Label}
		}
		return e.Faculty, nil
	}
	return nil, &NotLoadedError{edge: "faculty"}
}

// UserInformationsOrErr returns the UserInformations value or an error if the edge
// was not loaded in eager-loading.
func (e BranchEdges) UserInformationsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.UserInformations, nil
	}
	return nil, &NotLoadedError{edge: "user_informations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Branch) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // brname
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Branch) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // faculty_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Branch fields.
func (b *Branch) assignValues(values ...interface{}) error {
	if m, n := len(values), len(branch.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field brname", values[0])
	} else if value.Valid {
		b.Brname = value.String
	}
	values = values[1:]
	if len(values) == len(branch.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field faculty_id", value)
		} else if value.Valid {
			b.faculty_id = new(int)
			*b.faculty_id = int(value.Int64)
		}
	}
	return nil
}

// QueryFaculty queries the faculty edge of the Branch.
func (b *Branch) QueryFaculty() *FacultyQuery {
	return (&BranchClient{config: b.config}).QueryFaculty(b)
}

// QueryUserInformations queries the user_informations edge of the Branch.
func (b *Branch) QueryUserInformations() *UserQuery {
	return (&BranchClient{config: b.config}).QueryUserInformations(b)
}

// Update returns a builder for updating this Branch.
// Note that, you need to call Branch.Unwrap() before calling this method, if this Branch
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Branch) Update() *BranchUpdateOne {
	return (&BranchClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Branch) Unwrap() *Branch {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Branch is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Branch) String() string {
	var builder strings.Builder
	builder.WriteString("Branch(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", brname=")
	builder.WriteString(b.Brname)
	builder.WriteByte(')')
	return builder.String()
}

// Branches is a parsable slice of Branch.
type Branches []*Branch

func (b Branches) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
