// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/moominzie/user-record/ent/employee"
	"github.com/moominzie/user-record/ent/schema"
	"github.com/moominzie/user-record/ent/statust"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescEmployeename is the schema descriptor for employeename field.
	employeeDescEmployeename := employeeFields[0].Descriptor()
	// employee.EmployeenameValidator is a validator for the "employeename" field. It is called by the builders before save.
	employee.EmployeenameValidator = employeeDescEmployeename.Validators[0].(func(string) error)
	// employeeDescEmployeeemail is the schema descriptor for employeeemail field.
	employeeDescEmployeeemail := employeeFields[1].Descriptor()
	// employee.EmployeeemailValidator is a validator for the "employeeemail" field. It is called by the builders before save.
	employee.EmployeeemailValidator = employeeDescEmployeeemail.Validators[0].(func(string) error)
	// employeeDescPassword is the schema descriptor for password field.
	employeeDescPassword := employeeFields[2].Descriptor()
	// employee.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	employee.PasswordValidator = employeeDescPassword.Validators[0].(func(string) error)
	statustFields := schema.Statust{}.Fields()
	_ = statustFields
	// statustDescStatustname is the schema descriptor for statustname field.
	statustDescStatustname := statustFields[0].Descriptor()
	// statust.StatustnameValidator is a validator for the "statustname" field. It is called by the builders before save.
	statust.StatustnameValidator = statustDescStatustname.Validators[0].(func(string) error)
}