package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
       field.String("employeename").Unique().NotEmpty(),
	   field.String("employeeemail").NotEmpty(),
	   field.String("password").NotEmpty(),
   }
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Returninvoice.Type).StorageKey(edge.Column("employee_id")),
	}
}
