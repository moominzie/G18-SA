// Code generated by entc, DO NOT EDIT.

package branch

const (
	// Label holds the string label denoting the branch type in the database.
	Label = "branch"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBrname holds the string denoting the brname field in the database.
	FieldBrname = "brname"

	// EdgeFaculty holds the string denoting the faculty edge name in mutations.
	EdgeFaculty = "faculty"
	// EdgeUserInformations holds the string denoting the user_informations edge name in mutations.
	EdgeUserInformations = "user_informations"

	// Table holds the table name of the branch in the database.
	Table = "branches"
	// FacultyTable is the table the holds the faculty relation/edge.
	FacultyTable = "branches"
	// FacultyInverseTable is the table name for the Faculty entity.
	// It exists in this package in order to avoid circular dependency with the "faculty" package.
	FacultyInverseTable = "faculties"
	// FacultyColumn is the table column denoting the faculty relation/edge.
	FacultyColumn = "faculty_id"
	// UserInformationsTable is the table the holds the user_informations relation/edge.
	UserInformationsTable = "users"
	// UserInformationsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInformationsInverseTable = "users"
	// UserInformationsColumn is the table column denoting the user_informations relation/edge.
	UserInformationsColumn = "branch_id"
)

// Columns holds all SQL columns for branch fields.
var Columns = []string{
	FieldID,
	FieldBrname,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Branch type.
var ForeignKeys = []string{
	"faculty_id",
}