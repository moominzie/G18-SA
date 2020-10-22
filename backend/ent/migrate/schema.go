// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BranchesColumns holds the columns for the "branches" table.
	BranchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brname", Type: field.TypeString, Unique: true},
		{Name: "faculty_id", Type: field.TypeInt, Nullable: true},
	}
	// BranchesTable holds the schema information for the "branches" table.
	BranchesTable = &schema.Table{
		Name:       "branches",
		Columns:    BranchesColumns,
		PrimaryKey: []*schema.Column{BranchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "branches_faculties_branchs",
				Columns: []*schema.Column{BranchesColumns[2]},

				RefColumns: []*schema.Column{FacultiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BuildingsColumns holds the columns for the "buildings" table.
	BuildingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "buname", Type: field.TypeString, Unique: true},
	}
	// BuildingsTable holds the schema information for the "buildings" table.
	BuildingsTable = &schema.Table{
		Name:        "buildings",
		Columns:     BuildingsColumns,
		PrimaryKey:  []*schema.Column{BuildingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dname", Type: field.TypeString, Unique: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:        "devices",
		Columns:     DevicesColumns,
		PrimaryKey:  []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "employeename", Type: field.TypeString, Unique: true},
		{Name: "employeeemail", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:        "employees",
		Columns:     EmployeesColumns,
		PrimaryKey:  []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FacultiesColumns holds the columns for the "faculties" table.
	FacultiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fname", Type: field.TypeString, Unique: true},
	}
	// FacultiesTable holds the schema information for the "faculties" table.
	FacultiesTable = &schema.Table{
		Name:        "faculties",
		Columns:     FacultiesColumns,
		PrimaryKey:  []*schema.Column{FacultiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RepairInvoicesColumns holds the columns for the "repair_invoices" table.
	RepairInvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rename", Type: field.TypeString, Unique: true},
		{Name: "device_id", Type: field.TypeInt, Nullable: true},
		{Name: "statusr_id", Type: field.TypeInt, Nullable: true},
		{Name: "symptom_id", Type: field.TypeInt, Nullable: true},
		{Name: "repairinvoice_id", Type: field.TypeInt, Nullable: true},
	}
	// RepairInvoicesTable holds the schema information for the "repair_invoices" table.
	RepairInvoicesTable = &schema.Table{
		Name:       "repair_invoices",
		Columns:    RepairInvoicesColumns,
		PrimaryKey: []*schema.Column{RepairInvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "repair_invoices_devices_repair_information",
				Columns: []*schema.Column{RepairInvoicesColumns[2]},

				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "repair_invoices_status_rs_repair_information",
				Columns: []*schema.Column{RepairInvoicesColumns[3]},

				RefColumns: []*schema.Column{StatusRsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "repair_invoices_symptoms_repair_information",
				Columns: []*schema.Column{RepairInvoicesColumns[4]},

				RefColumns: []*schema.Column{SymptomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "repair_invoices_users_repairinvoice_informations",
				Columns: []*schema.Column{RepairInvoicesColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReturninvoicesColumns holds the columns for the "returninvoices" table.
	ReturninvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "returninvoice_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "statust_id", Type: field.TypeInt, Nullable: true},
	}
	// ReturninvoicesTable holds the schema information for the "returninvoices" table.
	ReturninvoicesTable = &schema.Table{
		Name:       "returninvoices",
		Columns:    ReturninvoicesColumns,
		PrimaryKey: []*schema.Column{ReturninvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "returninvoices_employees_employees",
				Columns: []*schema.Column{ReturninvoicesColumns[2]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "returninvoices_repair_invoices_returninvoice",
				Columns: []*schema.Column{ReturninvoicesColumns[3]},

				RefColumns: []*schema.Column{RepairInvoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "returninvoices_statusts_statusts",
				Columns: []*schema.Column{ReturninvoicesColumns[4]},

				RefColumns: []*schema.Column{StatustsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rname", Type: field.TypeString, Unique: true},
		{Name: "building_id", Type: field.TypeInt, Nullable: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "rooms_buildings_rooms",
				Columns: []*schema.Column{RoomsColumns[2]},

				RefColumns: []*schema.Column{BuildingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StatusRsColumns holds the columns for the "status_rs" table.
	StatusRsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sname", Type: field.TypeString, Unique: true},
	}
	// StatusRsTable holds the schema information for the "status_rs" table.
	StatusRsTable = &schema.Table{
		Name:        "status_rs",
		Columns:     StatusRsColumns,
		PrimaryKey:  []*schema.Column{StatusRsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatustsColumns holds the columns for the "statusts" table.
	StatustsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "statustname", Type: field.TypeString, Unique: true},
	}
	// StatustsTable holds the schema information for the "statusts" table.
	StatustsTable = &schema.Table{
		Name:        "statusts",
		Columns:     StatustsColumns,
		PrimaryKey:  []*schema.Column{StatustsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SymptomsColumns holds the columns for the "symptoms" table.
	SymptomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "syname", Type: field.TypeString, Unique: true},
	}
	// SymptomsTable holds the schema information for the "symptoms" table.
	SymptomsTable = &schema.Table{
		Name:        "symptoms",
		Columns:     SymptomsColumns,
		PrimaryKey:  []*schema.Column{SymptomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "personal_id", Type: field.TypeString, Unique: true},
		{Name: "personal_name", Type: field.TypeString},
		{Name: "branch_id", Type: field.TypeInt, Nullable: true},
		{Name: "building_id", Type: field.TypeInt, Nullable: true},
		{Name: "faculty_id", Type: field.TypeInt, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_branches_user_informations",
				Columns: []*schema.Column{UsersColumns[3]},

				RefColumns: []*schema.Column{BranchesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_buildings_user_informations",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{BuildingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_faculties_user_informations",
				Columns: []*schema.Column{UsersColumns[5]},

				RefColumns: []*schema.Column{FacultiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_rooms_user_informations",
				Columns: []*schema.Column{UsersColumns[6]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BranchesTable,
		BuildingsTable,
		DevicesTable,
		EmployeesTable,
		FacultiesTable,
		RepairInvoicesTable,
		ReturninvoicesTable,
		RoomsTable,
		StatusRsTable,
		StatustsTable,
		SymptomsTable,
		UsersTable,
	}
)

func init() {
	BranchesTable.ForeignKeys[0].RefTable = FacultiesTable
	RepairInvoicesTable.ForeignKeys[0].RefTable = DevicesTable
	RepairInvoicesTable.ForeignKeys[1].RefTable = StatusRsTable
	RepairInvoicesTable.ForeignKeys[2].RefTable = SymptomsTable
	RepairInvoicesTable.ForeignKeys[3].RefTable = UsersTable
	ReturninvoicesTable.ForeignKeys[0].RefTable = EmployeesTable
	ReturninvoicesTable.ForeignKeys[1].RefTable = RepairInvoicesTable
	ReturninvoicesTable.ForeignKeys[2].RefTable = StatustsTable
	RoomsTable.ForeignKeys[0].RefTable = BuildingsTable
	UsersTable.ForeignKeys[0].RefTable = BranchesTable
	UsersTable.ForeignKeys[1].RefTable = BuildingsTable
	UsersTable.ForeignKeys[2].RefTable = FacultiesTable
	UsersTable.ForeignKeys[3].RefTable = RoomsTable
}
