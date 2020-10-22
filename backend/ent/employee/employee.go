// Code generated by entc, DO NOT EDIT.

package employee

const (
	// Label holds the string label denoting the employee type in the database.
	Label = "employee"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmployeename holds the string denoting the employeename field in the database.
	FieldEmployeename = "employeename"
	// FieldEmployeeemail holds the string denoting the employeeemail field in the database.
	FieldEmployeeemail = "employeeemail"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeEmployees holds the string denoting the employees edge name in mutations.
	EdgeEmployees = "employees"
	// EdgeEmployeebill holds the string denoting the employeebill edge name in mutations.
	EdgeEmployeebill = "employeebill"

	// Table holds the table name of the employee in the database.
	Table = "employees"
	// EmployeesTable is the table the holds the employees relation/edge.
	EmployeesTable = "returninvoices"
	// EmployeesInverseTable is the table name for the Returninvoice entity.
	// It exists in this package in order to avoid circular dependency with the "returninvoice" package.
	EmployeesInverseTable = "returninvoices"
	// EmployeesColumn is the table column denoting the employees relation/edge.
	EmployeesColumn = "employee_id"
	// EmployeebillTable is the table the holds the employeebill relation/edge.
	EmployeebillTable = "bills"
	// EmployeebillInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	EmployeebillInverseTable = "bills"
	// EmployeebillColumn is the table column denoting the employeebill relation/edge.
	EmployeebillColumn = "employee_id"
)

// Columns holds all SQL columns for employee fields.
var Columns = []string{
	FieldID,
	FieldEmployeename,
	FieldEmployeeemail,
	FieldPassword,
}

var (
	// EmployeenameValidator is a validator for the "employeename" field. It is called by the builders before save.
	EmployeenameValidator func(string) error
	// EmployeeemailValidator is a validator for the "employeeemail" field. It is called by the builders before save.
	EmployeeemailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
